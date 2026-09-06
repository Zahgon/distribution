package proxy

import (
	"net/url"

	"github.com/distribution/distribution/v3/internal/client/auth"
	"github.com/distribution/distribution/v3/internal/client/auth/challenge"
)

const challengeHeader = "Docker-Distribution-Api-Version"

type userpass struct {
	username string
	password string
}

func (u userpass) Basic(_ *url.URL) (string, string) { _ = "STUB: not implemented"; return "", "" }

func (u userpass) RefreshToken(_ *url.URL, service string) string {
	_ = "STUB: not implemented"
	return ""
}

func (u userpass) SetRefreshToken(_ *url.URL, service, token string) {
	_ = "STUB: not implemented"
	return
}

type credentials struct {
	creds map[string]userpass
}

func (c credentials) Basic(u *url.URL) (string, string) { _ = "STUB: not implemented"; return "", "" }

func (c credentials) RefreshToken(u *url.URL, service string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c credentials) SetRefreshToken(u *url.URL, service, token string) {
	_ = "STUB: not implemented"
	return
}

func configureAuth(username, password, remoteURL string) (auth.CredentialStore, auth.CredentialStore, error) {
	_ = "STUB: not implemented"
	return *new(auth.CredentialStore), *new(auth.CredentialStore), nil
}

func getAuthURLs(remoteURL string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func realmAllowed(remote *url.URL, realm string) bool { _ = "STUB: not implemented"; return false }

func isLiteralOrLocal(host string) bool { _ = "STUB: not implemented"; return false }

func registrableDomain(host string) string { _ = "STUB: not implemented"; return "" }

func ping(manager challenge.Manager, endpoint, versionHeader string) error {
	_ = "STUB: not implemented"
	return nil
}
