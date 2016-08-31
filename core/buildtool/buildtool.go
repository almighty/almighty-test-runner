package buildtool

import (
    "path"
    "os"
)

// BuildTool represents basic information about the build tool which is in use
type BuildTool struct {
    Command, Path string
}

// Maven is a concrete type for Maven build system
type Maven struct {
    BuildTool
}

// InUse checks whether Maven is used as a build tool
func (m *Maven) InUse() bool {
    if _, err := os.Stat(path.Join(m.Path, "pom.xml")); err == nil {
        return true
    }
    return false
}

// DefaultCmd provides default command for Maven build
func (m *Maven) DefaultCmd() string {
    return "mvn clean install"
}

