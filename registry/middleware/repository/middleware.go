package middleware

import (
	"context"

	"github.com/distribution/distribution/v3"
)

type InitFunc func(ctx context.Context, repository distribution.Repository, options map[string]any) (distribution.Repository, error)

var middlewares map[string]InitFunc

func Register(name string, initFunc InitFunc) error { _ = "STUB: not implemented"; return nil }

func Get(ctx context.Context, name string, options map[string]any, repository distribution.Repository) (distribution.Repository, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Repository), nil
}
