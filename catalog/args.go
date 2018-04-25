package catalog

import "encoding/json"

// Args structure describes engine dependencies
type Args struct {
	// Engine id, case insensitive unique identifier, required only if the engine name is not unique
	ID string `json:"id,omitempty"`

	// Case sensitive official engine name, may not be unique
	Name string `json:"name"`

	Definitions *json.RawMessage `json:"definitions"`
}
