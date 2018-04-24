package catalog

// Args structure describes engine dependencies
type Args struct {
	// Engine id, case insensitive unique identifier, required only if the engine name is not unique
	ID string `json:"id"`

	// Case sensitive official engine name, may not be unique
	Name string `json:"name"`
}
