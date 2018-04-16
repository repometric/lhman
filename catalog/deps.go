package catalog

// Deps structure describes engine dependencies
type Deps struct {
	// Engine id, case insensitive unique identifier, required only if the engine name is not unique
	ID string `json:"id"`

	// Case sensitive official engine name, may not be unique
	Name string `json:"name"`

	// The complex object which describe several sub-schemas
	Dependencies [][]Requirement `json:"dependencies"`
}

// Requirement structure describes single dependency of engine
type Requirement struct {
	Manager string `json:"manager"`
	Package string `json:"package"`
	Version string `json:"version"`
	Engine  bool   `json:"engine"`
}
