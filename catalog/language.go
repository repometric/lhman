package catalog
type Language struct {
	//Array of language names
	Name []string `json:"enum"`
	//Array of language extensions
	Extensions []string `json:"extensions"`
}

