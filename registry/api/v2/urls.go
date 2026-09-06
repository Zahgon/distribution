package v2

import (
	"net/http"
	"net/url"

	"github.com/distribution/reference"
	"github.com/gorilla/mux"
)

type URLBuilder struct {
	root     *url.URL
	router   *mux.Router
	relative bool
}

func NewURLBuilder(root *url.URL, relative bool) *URLBuilder { _ = "STUB: not implemented"; return nil }

func NewURLBuilderFromString(root string, relative bool) (*URLBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewURLBuilderFromRequest(r *http.Request, relative bool) *URLBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (ub *URLBuilder) BuildBaseURL() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (ub *URLBuilder) BuildCatalogURL(values ...url.Values) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ub *URLBuilder) BuildTagsURL(name reference.Named, values ...url.Values) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ub *URLBuilder) BuildManifestURL(ref reference.Named) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ub *URLBuilder) BuildBlobURL(ref reference.Canonical) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ub *URLBuilder) BuildBlobUploadURL(name reference.Named, values ...url.Values) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ub *URLBuilder) BuildBlobUploadChunkURL(name reference.Named, uuid string, values ...url.Values) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ub *URLBuilder) cloneRoute(name string) clonedRoute {
	_ = "STUB: not implemented"
	return *new(clonedRoute)
}

type clonedRoute struct {
	*mux.Route
	root     *url.URL
	relative bool
}

func (cr clonedRoute) URL(pairs ...string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendValuesURL(u *url.URL, values ...url.Values) *url.URL {
	_ = "STUB: not implemented"
	return nil
}
