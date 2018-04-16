package install

import "github.com/repometric/lhman/catalog"

// Context describes all necessary parameters for installation
type Context struct {
	Engine      []catalog.Engine
	Project     string
	Environment string
	Version     []string
}
