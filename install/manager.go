package install

// Manager interface describes basic functionality for all managers in lhman
type Manager interface {
	Install(Context, Requirement) []ExecuteContext
	IsInstalled(Context, Requirement) bool
	InitManager() error
}
