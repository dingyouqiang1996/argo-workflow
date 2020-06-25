package auth

import (
	"errors"
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
	Token  Mode = "token"
)

func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso", "token":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true
	default:
		return errors.New("invalid mode")
	}
	return nil
}

const tokenPrefix = "Bearer token:"

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil
	}
	if strings.HasPrefix(authorisation, tokenPrefix) {
		return Token, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}
