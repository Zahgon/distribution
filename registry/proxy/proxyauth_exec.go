package proxy

import (
	"net/url"
	"sync"
	"time"

	"github.com/docker/docker-credential-helpers/client"
	credspkg "github.com/docker/docker-credential-helpers/credentials"

	"github.com/distribution/distribution/v3/configuration"
	"github.com/distribution/distribution/v3/internal/client/auth"
)

type execCredentials struct {
	m        sync.Mutex
	helper   client.ProgramFunc
	lifetime *time.Duration
	creds    *credspkg.Credentials
	expiry   time.Time
}

func (c *execCredentials) Basic(url *url.URL) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (c *execCredentials) RefreshToken(_ *url.URL, _ string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *execCredentials) SetRefreshToken(_ *url.URL, _, _ string) {
	_ = "STUB: not implemented"
	return
}

func configureExecAuth(cfg configuration.ExecConfig) (auth.CredentialStore, error) {
	_ = "STUB: not implemented"
	return *new(auth.CredentialStore), nil
}
