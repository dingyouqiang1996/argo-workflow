package util

import (
	"fmt"
	"github.com/argoproj/argo/errors"
	"github.com/argoproj/argo/util/retry"
	"github.com/prometheus/common/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

type Closer interface {
	Close() error
}

// Close is a convenience function to close a object that has a Close() method, ignoring any errors
// Used to satisfy errcheck lint
func Close(c Closer) {
	_ = c.Close()
}

// GetSecrets retrieves a secret value and memoizes the result
func GetSecrets(clientSet kubernetes.Interface, namespace, name, key string) ([]byte, error) {

	secretMaps := make(map[string][]byte)

	cachedKey := fmt.Sprintf("%s/%s/%s", namespace, name, key)

	secretsIf := clientSet.CoreV1().Secrets(namespace)
	var secret *apiv1.Secret
	var err error
	_ = wait.ExponentialBackoff(retry.DefaultRetry, func() (bool, error) {
		secret, err = secretsIf.Get(name, metav1.GetOptions{})
		if err != nil {
			log.Warnf("Failed to get secret '%s': %v", name, err)
			if !retry.IsRetryableKubeAPIError(err) {
				return false, err
			}
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return []byte{}, errors.InternalWrapError(err)
	}
	// memoize all keys in the secret since it's highly likely we will need to get a
	// subsequent key in the secret (e.g. username + password) and we can save an API call
	for k, v := range secret.Data {
		secretMaps[fmt.Sprintf("%s/%s/%s", namespace, name, k)] = v
	}
	val, ok := secretMaps[cachedKey]
	if !ok {
		return []byte{}, errors.Errorf(errors.CodeBadRequest, "secret '%s' does not have the key '%s'", name, key)
	}
	return val, nil
}
