package v2

import "github.com/distribution/distribution/v3/registry/api/errcode"

var (
	ErrorCodeDigestInvalid = errcode.ErrorCodeDigestInvalid

	ErrorCodeSizeInvalid = errcode.ErrorCodeSizeInvalid

	ErrorCodeRangeInvalid = errcode.ErrorCodeRangeInvalid

	ErrorCodeNameInvalid = errcode.ErrorCodeNameInvalid

	ErrorCodeTagInvalid = errcode.ErrorCodeTagInvalid

	ErrorCodeNameUnknown = errcode.ErrorCodeNameUnknown

	ErrorCodeManifestUnknown = errcode.ErrorCodeManifestUnknown

	ErrorCodeManifestInvalid = errcode.ErrorCodeManifestInvalid

	ErrorCodeManifestUnverified = errcode.ErrorCodeManifestUnverified

	ErrorCodeManifestBlobUnknown = errcode.ErrorCodeManifestBlobUnknown

	ErrorCodeBlobUnknown = errcode.ErrorCodeBlobUnknown

	ErrorCodeBlobUploadUnknown = errcode.ErrorCodeBlobUploadUnknown

	ErrorCodeBlobUploadInvalid = errcode.ErrorCodeBlobUploadInvalid

	ErrorCodePaginationNumberInvalid = errcode.ErrorCodePaginationNumberInvalid
)
