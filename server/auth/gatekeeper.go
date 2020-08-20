package auth

import (
	"context"
	"fmt"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/argoproj/argo/pkg/client/clientset/versioned"
	"github.com/argoproj/argo/server/auth/jws"
	"github.com/argoproj/argo/server/auth/jwt"
	"github.com/argoproj/argo/server/auth/sso"
	"github.com/argoproj/argo/util/kubeconfig"
)

type ContextKey string

const (
	WfKey       ContextKey = "versioned.Interface"
	KubeKey     ContextKey = "kubernetes.Interface"
	ClaimSetKey ContextKey = "jws.ClaimSet"
)

type Gatekeeper interface {
	Context(ctx context.Context) (context.Context, error)
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	StreamServerInterceptor() grpc.StreamServerInterceptor
}

type gatekeeper struct {
	Modes Modes
	// global clients, not to be used if there are better ones
	wfClient   versioned.Interface
	kubeClient kubernetes.Interface
	restConfig *rest.Config
	ssoIf      sso.Interface
	// The namespace the server is installed in.
	namespace string
}

func NewGatekeeper(modes Modes, wfClient versioned.Interface, kubeClient kubernetes.Interface, restConfig *rest.Config, ssoIf sso.Interface, namespace string) (Gatekeeper, error) {
	if len(modes) == 0 {
		return nil, fmt.Errorf("must specify at least one auth mode")
	}
	return &gatekeeper{modes, wfClient, kubeClient, restConfig, ssoIf, namespace}, nil
}

func (s *gatekeeper) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, err = s.Context(ctx)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (s *gatekeeper) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := s.Context(ss.Context())
		if err != nil {
			return err
		}
		wrapped := grpc_middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}

func (s *gatekeeper) Context(ctx context.Context) (context.Context, error) {
	wfClient, kubeClient, claimSet, err := s.getClients(ctx)
	if err != nil {
		return nil, err
	}
	return context.WithValue(context.WithValue(context.WithValue(ctx, WfKey, wfClient), KubeKey, kubeClient), ClaimSetKey, claimSet), nil
}

func GetWfClient(ctx context.Context) versioned.Interface {
	return ctx.Value(WfKey).(versioned.Interface)
}

func GetKubeClient(ctx context.Context) kubernetes.Interface {
	return ctx.Value(KubeKey).(kubernetes.Interface)
}

func GetClaimSet(ctx context.Context) *jws.ClaimSet {
	config, _ := ctx.Value(ClaimSetKey).(*jws.ClaimSet)
	return config
}

func getAuthHeader(md metadata.MD) string {
	// looks for the HTTP header `Authorization: Bearer ...`
	for _, t := range md.Get("authorization") {
		return t
	}
	// check the HTTP cookie
	for _, t := range md.Get("cookie") {
		header := http.Header{}
		header.Add("Cookie", t)
		request := http.Request{Header: header}
		token, err := request.Cookie("authorization")
		if err == nil {
			return token.Value
		}
	}
	return ""
}

func (s *gatekeeper) getClients(ctx context.Context) (versioned.Interface, kubernetes.Interface, *jws.ClaimSet, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	authorization := getAuthHeader(md)
	mode, err := GetMode(authorization)
	if err != nil {
		return nil, nil, nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if !s.Modes[mode] {
		return nil, nil, nil, status.Errorf(codes.Unauthenticated, "client auth-mode is %v, but that mode is disabled", mode)
	}
	switch mode {
	case Client:
		restConfig, wfClient, kubeClient, err := s.clientForAuthorization(authorization)
		if err != nil {
			return nil, nil, nil, status.Error(codes.Unauthenticated, err.Error())
		}
		claimSet, _ := jwt.ClaimSetFor(restConfig)
		return wfClient, kubeClient, claimSet, nil
	case Server:
		claimSet, _ := jwt.ClaimSetFor(s.restConfig)
		return s.wfClient, s.kubeClient, claimSet, nil
	case SSO:
		claimSet, err := s.ssoIf.Authorize(ctx, authorization)
		if err != nil {
			return nil, nil, nil, status.Error(codes.Unauthenticated, err.Error())
		}
		// RBAC for SSO maybe disabled, in that case both `serviceAccount` and `err` will be nil
		serviceAccount, err := s.ssoIf.GetServiceAccount(claimSet.Groups)
		if err != nil {
			return nil, nil, nil, status.Errorf(codes.Unauthenticated, "failed to get SSO RBAC service account ref: %v", err)
		} else if serviceAccount != nil {
			authorization, err := s.authorizationForServiceAccount(serviceAccount.Name)
			if err != nil {
				return nil, nil, nil, status.Error(codes.Unauthenticated, err.Error())
			}
			_, wfClient, kubeClient, err := s.clientForAuthorization(authorization)
			if err != nil {
				return nil, nil, nil, status.Error(codes.Unauthenticated, err.Error())
			}
			return wfClient, kubeClient, claimSet, nil
		} else {
			return s.wfClient, s.kubeClient, claimSet, nil
		}
	default:
		panic("this should never happen")
	}
}

func (s *gatekeeper) authorizationForServiceAccount(serviceAccountName string) (string, error) {
	serviceAccount, err := s.kubeClient.CoreV1().ServiceAccounts(s.namespace).Get(serviceAccountName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get SSO RBAC service account: %w", err)
	}
	if len(serviceAccount.Secrets) == 0 {
		return "", fmt.Errorf("expected at least one secret for SSO RBAC service account: %w", err)
	}
	secret, err := s.kubeClient.CoreV1().Secrets(s.namespace).Get(serviceAccount.Secrets[0].Name, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get SSO RBAC service account secret: %w", err)
	}
	return "Bearer " + string(secret.Data["token"]), nil
}

func (s *gatekeeper) clientForAuthorization(authorization string) (*rest.Config, versioned.Interface, kubernetes.Interface, error) {
	restConfig, err := kubeconfig.GetRestConfig(authorization)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create REST config: %w", err)
	}
	wfClient, err := versioned.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failure to create workflow client: %w", err)
	}
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failure to create kubernetes client: %w", err)
	}
	return restConfig, wfClient, kubeClient, nil
}
