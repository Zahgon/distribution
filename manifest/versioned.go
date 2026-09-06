package manifest

type Versioned struct {
	SchemaVersion int `json:"schemaVersion"`

	MediaType string `json:"mediaType,omitempty"`
}
