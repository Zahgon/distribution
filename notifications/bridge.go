package notifications

import (
	"net/http"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/reference"
	events "github.com/docker/go-events"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type bridge struct {
	ub                URLBuilder
	includeReferences bool
	actor             ActorRecord
	source            SourceRecord
	request           RequestRecord
	sink              events.Sink
}

var _ Listener = &bridge{}

type URLBuilder interface {
	BuildManifestURL(name reference.Named) (string, error)
	BuildBlobURL(ref reference.Canonical) (string, error)
}

func NewBridge(ub URLBuilder, source SourceRecord, actor ActorRecord, request RequestRecord, sink events.Sink, includeReferences bool) Listener {
	_ = "STUB: not implemented"
	return *new(Listener)
}

func NewRequestRecord(id string, r *http.Request) RequestRecord {
	_ = "STUB: not implemented"
	return *new(RequestRecord)
}

func (b *bridge) ManifestPushed(repo reference.Named, sm distribution.Manifest, options ...distribution.ManifestServiceOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) ManifestPulled(repo reference.Named, sm distribution.Manifest, options ...distribution.ManifestServiceOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) ManifestDeleted(repo reference.Named, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) BlobPushed(repo reference.Named, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) BlobPulled(repo reference.Named, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) BlobMounted(repo reference.Named, desc v1.Descriptor, fromRepo reference.Named) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) BlobDeleted(repo reference.Named, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) TagDeleted(repo reference.Named, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) RepoDeleted(repo reference.Named) error { _ = "STUB: not implemented"; return nil }

func (b *bridge) createManifestDeleteEventAndWrite(action string, repo reference.Named, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) createManifestEvent(action string, repo reference.Named, sm distribution.Manifest) (*Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bridge) createBlobDeleteEventAndWrite(action string, repo reference.Named, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) createBlobEventAndWrite(action string, repo reference.Named, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bridge) createBlobEvent(action string, repo reference.Named, desc v1.Descriptor) (*Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bridge) createEvent(action string) *Event { _ = "STUB: not implemented"; return nil }

func createEvent(action string) *Event { _ = "STUB: not implemented"; return nil }
