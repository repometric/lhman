package install

import (
	"os/exec"
	"regexp"
)

// NpmManager is a manager for nodejs
type NpmManager struct{}

// Install function of NpmManager installs npm dependencies
func (NpmManager) Install(context Context, requirement Requirement) string {
	globalFlag := ""
	if context.Environment == "global" {
		globalFlag = "-g"
	}
	packageFlag := requirement.Package
	if len(requirement.Version) > 0 {
		packageFlag += "@" + requirement.Version
	}
	_, err := Execute(context, requirement.Manager, globalFlag, packageFlag)

	return err
}

// IsInstalled function of NpmManager ckecks if npm dependency is installed
func (NpmManager) IsInstalled(context Context, requirement Requirement) bool {
	globalFlag := ""
	if context.Environment == "global" {
		globalFlag = "-g"
	}
	out, _ := Execute(context, requirement.Manager, globalFlag)
	reg, _ := regexp.Compile("\\s" + requirement.Package + "\\@" + requirement.Version)
	return len(reg.FindStringIndex(out)) != 0
}

// InitManager function checks if the package manager is available
func (m NpmManager) InitManager() error {
	_, error := exec.LookPath("npm")
	return error
}
