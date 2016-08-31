package buildtool

// Detector is the interface which holds the logic for detecting if a particular
// build system is in use and provides a default command to use this build tool with
type Detector interface {
    InUse(path string) bool
    DefaultCmd() string
}

