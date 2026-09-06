package v2

import (
	"sync"

	"github.com/gorilla/mux"
)

const (
	RouteNameBase            = "base"
	RouteNameManifest        = "manifest"
	RouteNameTags            = "tags"
	RouteNameBlob            = "blob"
	RouteNameBlobUpload      = "blob-upload"
	RouteNameBlobUploadChunk = "blob-upload-chunk"
	RouteNameCatalog         = "catalog"
)

var (
	baseRouter           *mux.Router
	createBaseRouterOnce sync.Once
)

func Router() *mux.Router { _ = "STUB: not implemented"; return nil }

func RouterWithPrefix(prefix string) *mux.Router { _ = "STUB: not implemented"; return nil }
