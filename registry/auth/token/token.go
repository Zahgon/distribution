package token

import (
	"crypto"
	"crypto/x509"
	"errors"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"

	"github.com/distribution/distribution/v3/registry/auth"
)

const (
	TokenSeparator = "."

	Leeway = 60 * time.Second
)

var signingAlgorithms = map[string]jose.SignatureAlgorithm{
	"EdDSA": jose.EdDSA,
	"HS256": jose.HS256,
	"HS384": jose.HS384,
	"HS512": jose.HS512,
	"RS256": jose.RS256,
	"RS384": jose.RS384,
	"RS512": jose.RS512,
	"ES256": jose.ES256,
	"ES384": jose.ES384,
	"ES512": jose.ES512,
	"PS256": jose.PS256,
	"PS384": jose.PS384,
	"PS512": jose.PS512,
}

var defaultSigningAlgorithms = []jose.SignatureAlgorithm{
	jose.EdDSA,
	jose.HS256,
	jose.HS384,
	jose.HS512,
	jose.RS256,
	jose.RS384,
	jose.RS512,
	jose.ES256,
	jose.ES384,
	jose.ES512,
	jose.PS256,
	jose.PS384,
	jose.PS512,
}

var (
	ErrMalformedToken = errors.New("malformed token")
	ErrInvalidToken   = errors.New("invalid token")
)

type ResourceActions struct {
	Type    string   `json:"type"`
	Class   string   `json:"class,omitempty"`
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

type ClaimSet struct {
	Issuer     string       `json:"iss"`
	Subject    string       `json:"sub"`
	Audience   AudienceList `json:"aud"`
	Expiration int64        `json:"exp"`
	NotBefore  int64        `json:"nbf"`
	IssuedAt   int64        `json:"iat"`
	JWTID      string       `json:"jti"`

	Access []*ResourceActions `json:"access"`
}

type Token struct {
	Raw string
	JWT *jwt.JSONWebToken
}

type VerifyOptions struct {
	TrustedIssuers    []string
	AcceptedAudiences []string
	Roots             *x509.CertPool
	TrustedKeys       map[string]crypto.PublicKey
}

func NewToken(rawToken string, signingAlgs []jose.SignatureAlgorithm) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Token) Verify(verifyOpts VerifyOptions) (*ClaimSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Token) VerifySigningKey(verifyOpts VerifyOptions) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

func verifyCertChain(header jose.Header, roots *x509.CertPool) (signingKey crypto.PublicKey, err error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

func verifyJWK(header jose.Header, verifyOpts VerifyOptions) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

func getCertPubKey(chains [][]*x509.Certificate) crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (c *ClaimSet) accessSet() accessSet { _ = "STUB: not implemented"; return *new(accessSet) }

func (c *ClaimSet) resources() []auth.Resource { _ = "STUB: not implemented"; return nil }
