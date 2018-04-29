package catalog
type Languages struct {
	Definitions Definitions `json:"definitions"`
}
type Definitions struct {
	Lang Lang `json:"language"`
}
type Lang struct {
	Language []Language `json:"oneOf"`
}
