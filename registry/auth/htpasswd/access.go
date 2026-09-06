package htpasswd

import (
	"net/http"
	"sync"
	"time"

	"github.com/distribution/distribution/v3/registry/auth"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := auth.Register("htpasswd", auth.InitFunc(newAccessController)); err != nil {
		logrus.Errorf("failed to register htpasswd auth: %v", err)
	}
}

type accessController struct {
	realm    string
	path     string
	modtime  time.Time
	mu       sync.Mutex
	htpasswd *htpasswd

	overrideDummyHash []byte
}

var _ auth.AccessController = &accessController{}

func newAccessController(options map[string]any) (auth.AccessController, error) {
	_ = "STUB: not implemented"
	return *new(auth.AccessController), nil
}

func (ac *accessController) Authorized(req *http.Request, accessRecords ...auth.Access) (*auth.Grant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type challenge struct {
	realm string
	err   error
}

var _ auth.Challenge = challenge{}

func (ch challenge) SetHeaders(r *http.Request, w http.ResponseWriter) {
	_ = "STUB: not implemented"
	return
}

func (ch challenge) Error() string { _ = "STUB: not implemented"; return "" }

func createHtpasswdFile(path string) error { _ = "STUB: not implemented"; return nil }
