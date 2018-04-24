package install

var managers = make(map[string]Manager)

func initManagers() {
	managers["platform"] = PlatformManager{}
	managers["npm"] = NpmManager{}
}

// Run function installs engines with given Context
func Run(context Context) []Result {
	if len(managers) == 0 {
		initManagers()
	}
	var results []Result
	for _, engine := range context.Engine {
		result := InitResult(context, engine)
		for branchIndex, dependencyBranch := range engine.Deps.Dependencies {
			for index, dependency := range dependencyBranch {
				requirement := InitRequirement(dependency)
				if requirement.Engine {
					if len(context.Version) > index {
						requirement.Version = context.Version[index]
					}
				}
				if manager, ok := managers[dependency.Manager]; ok {
					checkManager := manager.InitManager()
					if checkManager != nil {
						requirement.Errors = append(requirement.Errors, checkManager.Error())
						result.Success = false
					} else {
						if !manager.IsInstalled(context, requirement) {
							managerResult := manager.Install(context, requirement)
							for _, executeContext := range managerResult {
								_, stderr := Execute(executeContext)
								if len(stderr) != 0 {
									requirement.Errors = append(requirement.Errors, stderr)
									if requirement.Engine {
										result.Success = manager.IsInstalled(context, requirement)
									}
								}
							}
						}
					}
				} else {
					requirement.Errors = append(requirement.Errors, dependency.Manager+" is not supported")
					result.Success = false
				}
				result.Requirements = append(result.Requirements, requirement)
			}
			if result.Success {
				break
			} else if branchIndex < len(engine.Deps.Dependencies)-1 {
				result = InitResult(context, engine)
			}
		}
		results = append(results, result)
	}
	return results
}
