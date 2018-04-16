package catalog

// Meta structure describes engine metadata used in linterhub
type Meta struct {
	// Engine id, case insensitive unique identifier, required only if the engine name is not unique
	ID string `json:"id"`

	// Case sensitive official engine name, may not be unique
	Name string `json:"name"`

	// Official engine description, ideally "as is" and two sentences maximum
	Description string `json:"description"`

	// Url to engine homepage (official website, repository or documentation)
	URL string `json:"url"`

	// List of supported languages; if not set then assumed that all languages are supported
	Languages []string `json:"languages"`

	// List of filename masks that are supported by the engine, normally it's just extensions;
	// if not set then assumed that all extensions are supported for chosen language(s);
	// it's also allowed to use language name there
	Extensions []string `json:"extensions"`

	// List of filename masks that could be treated as engine config file, normally it's a filename with extension;
	// if not set then assumed that engine has no config files
	Configs []string `json:"configs"`

	// The license name of the original linter;
	// if linter has a custom license then this property should be equal to `custom` and the agreement property set accordingly
	License string `json:"license"`

	// Url or path to the custom license content;
	// if not set then assumed that custom license is located at `license.txt` in the engine folder
	Agreement string `json:"agreement"`
}
