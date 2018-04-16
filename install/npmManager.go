package install

import (
	"bytes"
	"os/exec"
	"regexp"
)

// NpmManager is a manager for nodejs
type NpmManager struct{}

// Install function of NpmManager installs npm dependencies
func (m NpmManager) Install(c Context, r Requirement) string {
	globalFlag := ""
	if c.Environment == "global" {
		globalFlag = "-g"
	}
	packageFlag := r.Package
	if len(r.Version) > 0 {
		packageFlag += "@" + r.Version
	}
	cmd := exec.Command(r.Manager, "install", globalFlag, packageFlag)
	if len(c.Project) > 0 {
		cmd.Dir = c.Project
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "npm manager crashed with: " + err.Error()
	}

	return string(stderr.Bytes())
}

// IsInstalled function of NpmManager ckecks if npm dependency is installed
func (m NpmManager) IsInstalled(c Context, r Requirement) bool {
	globalFlag := ""
	if c.Environment == "global" {
		globalFlag = "-g"
	}
	cmd := exec.Command(r.Manager, "list", globalFlag)
	if len(c.Project) > 0 {
		cmd.Dir = c.Project
	}
	out, _ := cmd.Output()
	reg, _ := regexp.Compile("\\s" + r.Package + "\\@" + r.Version)
	return len(reg.FindStringIndex(string(out[:bytes.IndexByte(out, 0)]))) != 0
}

// InitManager function checks if the package manager is available
func (m NpmManager) InitManager() error {
	_, error := exec.LookPath("npm")
	return error
}
