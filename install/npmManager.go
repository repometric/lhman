package install

import (
	"os/exec"
	"regexp"
)

// NpmManager is a manager for nodejs
type NpmManager struct{}

// Install function of NpmManager installs npm dependencies
func (NpmManager) Install(context Context, requirement Requirement) []ExecuteContext {
	args := []string{"install"}
	if context.Environment == "global" {
		args = append(args, "-g")
	}
	packageFlag := requirement.Package
	if len(requirement.Version) > 0 {
		packageFlag += "@" + requirement.Version
	}
	args = append(args, packageFlag)

	return []ExecuteContext{
		ExecuteContext{
			Binary:           requirement.Manager,
			WorkingDirectory: context.Folder,
			Args:             args,
		},
	}
}

// IsInstalled function of NpmManager ckecks if npm dependency is installed
func (NpmManager) IsInstalled(context Context, requirement Requirement) bool {
	args := []string{"list"}
	if context.Environment == "global" {
		args = append(args, "-g")
	}
	out, _ := Execute(
		ExecuteContext{
			Binary:           requirement.Manager,
			WorkingDirectory: context.Folder,
			Args:             args,
		},
	)
	reg, _ := regexp.Compile("\\s" + requirement.Package + "\\@" + requirement.Version)
	return len(reg.FindStringIndex(out)) != 0
}

// InitManager function checks if the package manager is available
func (m NpmManager) InitManager() error {
	_, error := exec.LookPath("npm")
	return error
}
