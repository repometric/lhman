package install

// Manager interface describes basic functionality for all managers in lhman
type Manager interface {
	Install(Context, Requirement) string
	IsInstalled(Context, Requirement) bool
	InitManager() error
}
