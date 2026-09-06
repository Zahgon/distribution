package errcode

import (
	"net/http"
	"sync"
)

var (
	errorCodeToDescriptors = map[ErrorCode]ErrorDescriptor{}
	idToDescriptors        = map[string]ErrorDescriptor{}
	groupToDescriptors     = map[string][]ErrorDescriptor{}
)

var (
	ErrorCodeUnknown = register("errcode", ErrorDescriptor{
		Value:   "UNKNOWN",
		Message: "unknown error",
		Description: `Generic error returned when the error does not have an
			                                            API classification.`,
		HTTPStatusCode: http.StatusInternalServerError,
	})

	ErrorCodeUnsupported = register("errcode", ErrorDescriptor{
		Value:   "UNSUPPORTED",
		Message: "The operation is unsupported.",
		Description: `The operation was unsupported due to a missing
		implementation or invalid set of parameters.`,
		HTTPStatusCode: http.StatusMethodNotAllowed,
	})

	ErrorCodeUnauthorized = register("errcode", ErrorDescriptor{
		Value:   "UNAUTHORIZED",
		Message: "authentication required",
		Description: `The access controller was unable to authenticate
		the client. Often this will be accompanied by a
		Www-Authenticate HTTP response header indicating how to
		authenticate.`,
		HTTPStatusCode: http.StatusUnauthorized,
	})

	ErrorCodeDenied = register("errcode", ErrorDescriptor{
		Value:   "DENIED",
		Message: "requested access to the resource is denied",
		Description: `The access controller denied access for the
		operation on a resource.`,
		HTTPStatusCode: http.StatusForbidden,
	})

	ErrorCodeUnavailable = register("errcode", ErrorDescriptor{
		Value:          "UNAVAILABLE",
		Message:        "service unavailable",
		Description:    "Returned when a service is not available",
		HTTPStatusCode: http.StatusServiceUnavailable,
	})

	ErrorCodeTooManyRequests = register("errcode", ErrorDescriptor{
		Value:   "TOOMANYREQUESTS",
		Message: "too many requests",
		Description: `Returned when a client attempts to contact a
		service too many times`,
		HTTPStatusCode: http.StatusTooManyRequests,
	})
)

const errGroup = "registry.api.v2"

var (
	ErrorCodeDigestInvalid = register(errGroup, ErrorDescriptor{
		Value:   "DIGEST_INVALID",
		Message: "provided digest did not match uploaded content",
		Description: `When a blob is uploaded, the registry will check that
		the content matches the digest provided by the client. The error may
		include a detail structure with the key "digest", including the
		invalid digest string. This error may also be returned when a manifest
		includes an invalid layer digest.`,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeSizeInvalid = register(errGroup, ErrorDescriptor{
		Value:   "SIZE_INVALID",
		Message: "provided length did not match content length",
		Description: `When a layer is uploaded, the provided size will be
		checked against the uploaded content. If they do not match, this error
		will be returned.`,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeRangeInvalid = register(errGroup, ErrorDescriptor{
		Value:   "RANGE_INVALID",
		Message: "invalid content range",
		Description: `When a layer is uploaded, the provided range is checked
		against the uploaded chunk. This error is returned if the range is
		out of order.`,
		HTTPStatusCode: http.StatusRequestedRangeNotSatisfiable,
	})

	ErrorCodeNameInvalid = register(errGroup, ErrorDescriptor{
		Value:   "NAME_INVALID",
		Message: "invalid repository name",
		Description: `Invalid repository name encountered either during
		manifest validation or any API operation.`,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeTagInvalid = register(errGroup, ErrorDescriptor{
		Value:   "TAG_INVALID",
		Message: "manifest tag did not match URI",
		Description: `During a manifest upload, if the tag in the manifest
		does not match the uri tag, this error will be returned.`,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeNameUnknown = register(errGroup, ErrorDescriptor{
		Value:   "NAME_UNKNOWN",
		Message: "repository name not known to registry",
		Description: `This is returned if the name used during an operation is
		unknown to the registry.`,
		HTTPStatusCode: http.StatusNotFound,
	})

	ErrorCodeManifestUnknown = register(errGroup, ErrorDescriptor{
		Value:   "MANIFEST_UNKNOWN",
		Message: "manifest unknown",
		Description: `This error is returned when the manifest, identified by
		name and tag is unknown to the repository.`,
		HTTPStatusCode: http.StatusNotFound,
	})

	ErrorCodeManifestInvalid = register(errGroup, ErrorDescriptor{
		Value:   "MANIFEST_INVALID",
		Message: "manifest invalid",
		Description: `During upload, manifests undergo several checks ensuring
		validity. If those checks fail, this error may be returned, unless a
		more specific error is included. The detail will contain information
		the failed validation.`,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeManifestUnverified = register(errGroup, ErrorDescriptor{
		Value:   "MANIFEST_UNVERIFIED",
		Message: "manifest failed signature verification",
		Description: `During manifest upload, if the manifest fails signature
		verification, this error will be returned.`,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeManifestBlobUnknown = register(errGroup, ErrorDescriptor{
		Value:   "MANIFEST_BLOB_UNKNOWN",
		Message: "blob unknown to registry",
		Description: `This error may be returned when a manifest blob is 
		unknown to the registry.`,
		HTTPStatusCode: http.StatusBadRequest,
	})

	ErrorCodeBlobUnknown = register(errGroup, ErrorDescriptor{
		Value:   "BLOB_UNKNOWN",
		Message: "blob unknown to registry",
		Description: `This error may be returned when a blob is unknown to the
		registry in a specified repository. This can be returned with a
		standard get or if a manifest references an unknown layer during
		upload.`,
		HTTPStatusCode: http.StatusNotFound,
	})

	ErrorCodeBlobUploadUnknown = register(errGroup, ErrorDescriptor{
		Value:   "BLOB_UPLOAD_UNKNOWN",
		Message: "blob upload unknown to registry",
		Description: `If a blob upload has been cancelled or was never
		started, this error code may be returned.`,
		HTTPStatusCode: http.StatusNotFound,
	})

	ErrorCodeBlobUploadInvalid = register(errGroup, ErrorDescriptor{
		Value:   "BLOB_UPLOAD_INVALID",
		Message: "blob upload invalid",
		Description: `The blob upload encountered an error and can no
		longer proceed.`,
		HTTPStatusCode: http.StatusNotFound,
	})

	ErrorCodePaginationNumberInvalid = register(errGroup, ErrorDescriptor{
		Value:   "PAGINATION_NUMBER_INVALID",
		Message: "invalid number of results requested",
		Description: `Returned when the "n" parameter (number of results
		to return) is not an integer, "n" is negative or "n" is bigger than
		the maximum allowed.`,
		HTTPStatusCode: http.StatusBadRequest,
	})
)

var (
	nextCode     = 1000
	registerLock sync.Mutex
)

func Register(group string, descriptor ErrorDescriptor) ErrorCode {
	_ = "STUB: not implemented"
	return *new(ErrorCode)
}

func register(group string, descriptor ErrorDescriptor) ErrorCode {
	_ = "STUB: not implemented"
	return *new(ErrorCode)
}

type byValue []ErrorDescriptor

func (a byValue) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a byValue) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a byValue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func GetGroupNames() []string { _ = "STUB: not implemented"; return nil }

func GetErrorCodeGroup(name string) []ErrorDescriptor { _ = "STUB: not implemented"; return nil }

func GetErrorAllDescriptors() []ErrorDescriptor { _ = "STUB: not implemented"; return nil }
