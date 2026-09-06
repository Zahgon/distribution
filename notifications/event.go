package notifications

import (
	"fmt"
	"time"

	events "github.com/docker/go-events"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	EventActionPull   = "pull"
	EventActionPush   = "push"
	EventActionMount  = "mount"
	EventActionDelete = "delete"
)

const (
	EventsMediaType = "application/vnd.docker.distribution.events.v2+json"

	layerMediaType = "application/vnd.docker.container.image.rootfs.diff+x-gtar"
)

type Envelope struct {
	Events []events.Event `json:"events,omitempty"`
}

type Event struct {
	ID string `json:"id,omitempty"`

	Timestamp time.Time `json:"timestamp"`

	Action string `json:"action,omitempty"`

	Target struct {
		v1.Descriptor

		Length int64 `json:"length,omitempty"`

		Repository string `json:"repository,omitempty"`

		FromRepository string `json:"fromRepository,omitempty"`

		URL string `json:"url,omitempty"`

		Tag string `json:"tag,omitempty"`

		References []v1.Descriptor `json:"references,omitempty"`
	} `json:"target"`

	Request RequestRecord `json:"request"`

	Actor ActorRecord `json:"actor"`

	Source SourceRecord `json:"source"`
}

type ActorRecord struct {
	Name string `json:"name,omitempty"`
}

type RequestRecord struct {
	ID string `json:"id"`

	Addr string `json:"addr,omitempty"`

	Host string `json:"host,omitempty"`

	Method string `json:"method"`

	UserAgent string `json:"useragent"`
}

type SourceRecord struct {
	Addr string `json:"addr,omitempty"`

	InstanceID string `json:"instanceID,omitempty"`
}

var ErrSinkClosed = fmt.Errorf("sink: closed")
