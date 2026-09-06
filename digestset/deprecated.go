package digestset

import (
	"github.com/opencontainers/go-digest"
	"github.com/opencontainers/go-digest/digestset"
)

var ErrDigestNotFound = digestset.ErrDigestNotFound

var ErrDigestAmbiguous = digestset.ErrDigestAmbiguous

type Set = digestset.Set

func NewSet() *digestset.Set { _ = "STUB: not implemented"; return nil }

func ShortCodeTable(dst *digestset.Set, length int) map[digest.Digest]string {
	_ = "STUB: not implemented"
	return nil
}
