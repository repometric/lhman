package install

import "os/exec"

// PlatformManager is a manager for checking and installing global packages like other package managers
type PlatformManager struct{}

// Install function of PlatformManager installs global Requirements
func (m PlatformManager) Install(c Context, r Requirement) string {
	return "lhman cant install global packages like `" + r.Package + "`"
}

// IsInstalled function of PlatformManager ckecks global Requirements
func (m PlatformManager) IsInstalled(c Context, r Requirement) bool {
	_, lookErr := exec.LookPath(r.Package) // TODO: check version of package
	return lookErr != nil
}

// InitManager function checks if the package manager is available
func (m PlatformManager) InitManager() error {
	return nil
}
