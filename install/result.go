package install

import "github.com/repometric/lhman/catalog"

// Result describes install output for engine
type Result struct {
	ID           string        `json:"id"`
	Environment  string        `json:"environment"`
	Success      bool          `json:"success"`
	Requirements []Requirement `json:"requirements"`
}

// InitResult function creates basic instance of Result for Engine
func InitResult(c Context, e catalog.Engine) Result {
	id := e.Meta.ID
	if id == "" {
		id = e.Meta.Name
	}
	return Result{
		Environment:  c.Environment,
		ID:           id,
		Success:      true,
		Requirements: make([]Requirement, 0),
	}
}

// Requirement describes install result of engine's dependency
type Requirement struct {
	Manager string `json:"manager"`
	Package string `json:"package"`
	Version string `json:"version"`
	Engine  bool   `json:"engine"`
	Errors  string `json:"errors"`
}

// InitRequirement function creates basic instance of Requirement
func InitRequirement(req catalog.Requirement) Requirement {
	return Requirement{
		Manager: req.Manager,
		Package: req.Package,
		Version: req.Version,
		Engine:  req.Engine,
	}
}
