package catalog

type Engine struct {
	Meta Meta
}

type Meta struct {
	Name string
	Description string
	Url string
	Languages []string
	Configs []string
	License string
}