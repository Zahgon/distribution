package auth

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidCredential = errors.New("invalid authorization credential")

	ErrAuthenticationFailure = errors.New("authentication failure")
)

type InitFunc func(options map[string]any) (AccessController, error)

var accessControllers map[string]InitFunc

func init() {
	accessControllers = make(map[string]InitFunc)
}

type UserInfo struct {
	Name string
}

type Resource struct {
	Type  string
	Class string
	Name  string
}

type Access struct {
	Resource
	Action string
}

type Grant struct {
	User      UserInfo
	Resources []Resource
}

type Challenge interface {
	error

	SetHeaders(r *http.Request, w http.ResponseWriter)
}

type AccessController interface {
	Authorized(r *http.Request, access ...Access) (*Grant, error)
}

type CredentialAuthenticator interface {
	AuthenticateUser(username, password string) error
}

func Register(name string, initFunc InitFunc) error { _ = "STUB: not implemented"; return nil }

func GetAccessController(name string, options map[string]any) (AccessController, error) {
	_ = "STUB: not implemented"
	return *new(AccessController), nil
}
