package install

var managers = make(map[string]Manager)

func initManagers() {
	managers["platform"] = PlatformManager{}
	managers["npm"] = NpmManager{}
}

// Run function installs engines witg given Context
func Run(c Context) []Result {
	if len(managers) == 0 {
		initManagers()
	}
	var results []Result
	for _, engine := range c.Engine {
		result := InitResult(c, engine)
		var dependencyBranch = engine.Deps.Dependencies[0] // TODO
		for index, dependency := range dependencyBranch {
			requirement := InitRequirement(dependency)
			if requirement.Engine {
				if len(c.Version) > index {
					requirement.Version = c.Version[index]
				}
			}
			checkManager := managers[dependency.Manager].InitManager()
			if checkManager != nil {
				requirement.Errors = checkManager.Error()
				result.Success = false
			} else {
				if !managers[dependency.Manager].IsInstalled(c, requirement) {
					managerResult := managers[dependency.Manager].Install(c, requirement)
					if len(managerResult) != 0 {
						requirement.Errors = managerResult
						result.Success = false
					}
				}
			}
			result.Requirements = append(result.Requirements, requirement)

		}
		results = append(results, result)
	}
	return results
}
