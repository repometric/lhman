package install

// ExecuteContext describes params necessary for command execution
type ExecuteContext struct {
	Binary           string
	WorkingDirectory string
	Args             []string
}
