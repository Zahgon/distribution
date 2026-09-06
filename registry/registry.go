package registry

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/acme"

	"github.com/distribution/distribution/v3/configuration"
	"github.com/distribution/distribution/v3/internal/dcontext"
	"github.com/distribution/distribution/v3/registry/handlers"
	"github.com/distribution/distribution/v3/version"
)

var cipherSuites = map[string]uint16{

	"TLS_RSA_WITH_3DES_EDE_CBC_SHA":                 tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_RSA_WITH_AES_128_CBC_SHA":                  tls.TLS_RSA_WITH_AES_128_CBC_SHA,
	"TLS_RSA_WITH_AES_256_CBC_SHA":                  tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	"TLS_RSA_WITH_AES_128_GCM_SHA256":               tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_RSA_WITH_AES_256_GCM_SHA384":               tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":          tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":          tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":           tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":            tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":         tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256":       tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":         tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384":       tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256":   tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
	"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256": tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,

	"TLS_AES_128_GCM_SHA256":       tls.TLS_AES_128_GCM_SHA256,
	"TLS_AES_256_GCM_SHA384":       tls.TLS_AES_256_GCM_SHA384,
	"TLS_CHACHA20_POLY1305_SHA256": tls.TLS_CHACHA20_POLY1305_SHA256,
}

var defaultCipherSuites = []uint16{
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_AES_128_GCM_SHA256,
	tls.TLS_CHACHA20_POLY1305_SHA256,
	tls.TLS_AES_256_GCM_SHA384,
}

const defaultTLSVersionStr = "tls1.2"

var tlsVersions = map[string]uint16{
	"tls1.2": tls.VersionTLS12,
	"tls1.3": tls.VersionTLS13,
}

var tlsClientAuth = map[string]tls.ClientAuthType{
	configuration.ClientAuthRequestClientCert:          tls.RequestClientCert,
	configuration.ClientAuthRequireAnyClientCert:       tls.RequireAnyClientCert,
	configuration.ClientAuthVerifyClientCertIfGiven:    tls.VerifyClientCertIfGiven,
	configuration.ClientAuthRequireAndVerifyClientCert: tls.RequireAndVerifyClientCert,
}

const defaultLogFormatter = "text"

type HandlerFunc func(config *configuration.Configuration, handler http.Handler) http.Handler

var handlerMiddlewares []HandlerFunc

func RegisterHandler(handlerFunc HandlerFunc) { _ = "STUB: not implemented"; return }

var ServeCmd = &cobra.Command{
	Use:   "serve <config>",
	Short: "`serve` stores and distributes Docker images",
	Long:  "`serve` stores and distributes Docker images.",
	Run: func(cmd *cobra.Command, args []string) {

		ctx := dcontext.WithVersion(dcontext.Background(), version.Version())

		config, err := resolveConfiguration(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)

			cmd.Usage()
			os.Exit(1)
		}
		registry, err := NewRegistry(ctx, config)
		if err != nil {
			logrus.Fatalln(err)
		}

		configureDebugServer(config)

		if err = registry.ListenAndServe(); err != nil {
			logrus.Fatalln(err)
		}
	},
}

type Registry struct {
	config *configuration.Configuration
	app    *handlers.App
	server *http.Server
	quit   chan os.Signal
}

func NewRegistry(ctx context.Context, config *configuration.Configuration) (*Registry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func otelHandler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func getCipherSuites(names []string) ([]uint16, error) { _ = "STUB: not implemented"; return nil, nil }

func getCipherSuiteNames(ids []uint16) []string { _ = "STUB: not implemented"; return nil }

func setDirectoryURL(directoryurl string) *acme.Client { _ = "STUB: not implemented"; return nil }

func (registry *Registry) ListenAndServe() error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck // FIXME(thaJeztah): ignore SA1019: ac.(*accessController).rootCerts.Subjects has been deprecated since Go 1.18: if s was returned by SystemCertPool, Subjects will not include the system roots. (staticcheck)

func (registry *Registry) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func configureDebugServer(config *configuration.Configuration) { _ = "STUB: not implemented"; return }

func configurePrometheus(config *configuration.Configuration) { _ = "STUB: not implemented"; return }

func configureLogging(ctx context.Context, config *configuration.Configuration) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func logLevel(level configuration.Loglevel) logrus.Level {
	_ = "STUB: not implemented"
	return *new(logrus.Level)
}

func panicHandler(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func alive(path string, handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func resolveConfiguration(args []string) (*configuration.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func nextProtos(config *configuration.Configuration) []string {
	_ = "STUB: not implemented"
	return nil
}

func serverProtocols(config *configuration.Configuration) *http.Protocols {
	_ = "STUB: not implemented"
	return nil
}
