package proxy

import (
	"context"

	"github.com/distribution/distribution/v3"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type proxyTagService struct {
	localTags      distribution.TagService
	remoteTags     distribution.TagService
	authChallenger authChallenger
}

var _ distribution.TagService = proxyTagService{}

func (pt proxyTagService) Get(ctx context.Context, tag string) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (pt proxyTagService) Tag(ctx context.Context, tag string, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (pt proxyTagService) Untag(ctx context.Context, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pt proxyTagService) All(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pt proxyTagService) Lookup(ctx context.Context, digest v1.Descriptor) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pt proxyTagService) List(ctx context.Context, limit int, last string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
