package htpasswd

import (
	"context"
	"io"
)

const dummyBcryptHash = "$2a$05$AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

type htpasswd struct {
	entries map[string][]byte

	dummyHash []byte
}

func newHTPasswd(rd io.Reader, dummyHash []byte) (*htpasswd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (htpasswd *htpasswd) authenticateUser(ctx context.Context, username string, password string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseHTPasswd(rd io.Reader) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
