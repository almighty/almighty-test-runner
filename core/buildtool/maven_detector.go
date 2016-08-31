package buildtool

import (
    "github.com/almighty/almighty-test-runner/core/configuration"
)

// Create is a factory method which creates Maven instance based on the
// Configuration and path of the project
func Create(path string, configuration configuration.TestRunnerConfiguration) *Maven {
    return &Maven{BuildTool{Command: configuration.Command, Path: path}}
}


