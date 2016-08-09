package buildtool

import (
    "github.com/almighty/almighty-test-runner/core/configuration"
)

func Create(path string, configuration configuration.TestRunnerConfiguration) *Maven {
    return &Maven{BuildTool{Command: configuration.Command, Path: path}}
}


