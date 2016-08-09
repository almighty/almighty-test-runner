package buildtool

import (
    "path"
    "os"
)

type BuildTool struct {
    Command, Path string
}

type Maven struct {
    BuildTool
}

func (m *Maven) InUse() bool {
    if _, err := os.Stat(path.Join(m.Path, "pom.xml")); err == nil {
        return true
    }
    return false
}

func (m *Maven) DefaultCmd() string {
    return "mvn clean install"
}

