package catalog

// Engine structure
type Engine struct {
	Meta Meta `json:"meta"`
	Deps Deps `json:"deps"`
}
