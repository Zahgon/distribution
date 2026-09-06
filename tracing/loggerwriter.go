package tracing

import "github.com/distribution/distribution/v3/internal/dcontext"

type loggerWriter struct {
	logger dcontext.Logger
}

func (lw *loggerWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (lw *loggerWriter) Handle(err error) { _ = "STUB: not implemented"; return }
