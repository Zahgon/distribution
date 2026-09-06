package distribution

import (
	"errors"

	"github.com/opencontainers/go-digest"
)

var ErrAccessDenied = errors.New("access denied")

var ErrManifestNotModified = errors.New("manifest not modified")

var ErrUnsupported = errors.New("operation unsupported")

var ErrSchemaV1Unsupported = errors.New("manifest schema v1 unsupported")

type ErrTagUnknown struct {
	Tag string
}

func (err ErrTagUnknown) Error() string { _ = "STUB: not implemented"; return "" }

type ErrRepositoryUnknown struct {
	Name string
}

func (err ErrRepositoryUnknown) Error() string { _ = "STUB: not implemented"; return "" }

type ErrRepositoryNameInvalid struct {
	Name   string
	Reason error
}

func (err ErrRepositoryNameInvalid) Error() string { _ = "STUB: not implemented"; return "" }

type ErrManifestUnknown struct {
	Name string
	Tag  string
}

func (err ErrManifestUnknown) Error() string { _ = "STUB: not implemented"; return "" }

type ErrManifestUnknownRevision struct {
	Name     string
	Revision digest.Digest
}

func (err ErrManifestUnknownRevision) Error() string { _ = "STUB: not implemented"; return "" }

type ErrManifestUnverified struct{}

func (ErrManifestUnverified) Error() string { _ = "STUB: not implemented"; return "" }

type ErrManifestVerification []error

func (errs ErrManifestVerification) Error() string { _ = "STUB: not implemented"; return "" }

type ErrManifestBlobUnknown struct {
	Digest digest.Digest
}

func (err ErrManifestBlobUnknown) Error() string { _ = "STUB: not implemented"; return "" }

type ErrManifestNameInvalid struct {
	Name   string
	Reason error
}

func (err ErrManifestNameInvalid) Error() string { _ = "STUB: not implemented"; return "" }
