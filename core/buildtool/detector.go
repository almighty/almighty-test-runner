package buildtool

type Detector interface {
    InUse(path string) bool
    DefaultCmd() string
}

