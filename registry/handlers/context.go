package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/api/errcode"
	v2 "github.com/distribution/distribution/v3/registry/api/v2"
	"github.com/distribution/distribution/v3/registry/auth"
	"github.com/opencontainers/go-digest"
)

type Context struct {
	*App
	context.Context

	Repository distribution.Repository

	RepositoryRemover distribution.RepositoryRemover

	Errors errcode.Errors

	urlBuilder *v2.URLBuilder
}

func (ctx *Context) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }

func getName(ctx context.Context) (name string) { _ = "STUB: not implemented"; return "" }

func getReference(ctx context.Context) (reference string) { _ = "STUB: not implemented"; return "" }

var errDigestNotAvailable = fmt.Errorf("digest not available in context")

func getDigest(ctx context.Context) (dgst digest.Digest, err error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

func getUploadUUID(ctx context.Context) (uuid string) { _ = "STUB: not implemented"; return "" }

const (
	userKey = "auth.user"

	userNameKey = "auth.user.name"
)

func getUserName(ctx context.Context, r *http.Request) string { _ = "STUB: not implemented"; return "" }

func withUser(ctx context.Context, user auth.UserInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type userInfoContext struct {
	context.Context
	user auth.UserInfo
}

func (uic userInfoContext) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }

func withResources(ctx context.Context, resources []auth.Resource) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type resourceContext struct {
	context.Context
	resources []auth.Resource
}

type resourceKey struct{}

func (rc resourceContext) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }

func authorizedResources(ctx context.Context) []auth.Resource {
	_ = "STUB: not implemented"
	return nil
}
