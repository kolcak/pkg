package build

import (
	"os"
	"path/filepath"
	"strings"
)

// On build `-ldflags` values
var (
	xEnv      = "ENV"
	xProject  = "project"
	xVersion  = "v0.0.0"
	xRevision = "v000"
	xRelease  = "0"
)

var build *Build = &Build{
	Env:      strings.ToUpper(xEnv),
	Project:  strings.ToLower(xProject),
	Version:  xVersion,
	Revision: xRevision,
	Release:  xRelease,
}

type Build struct {
	Env      string
	Project  string
	Version  string
	Revision string
	Release  string
}

func init() {
	project, err := os.Executable()
	if err == nil {
		build.Project = strings.ToLower(filepath.Base(project))
	}
}

func Info() *Build {
	return build
}
