package mockup_test

import (
	"fmt"
	"strings"

	"github.com/kolcak/pkg/src/options"
)

type TestOptions struct {
	Setup string
}

func (t *TestOptions) Default() interface{} {
	return &TestOptions{
		Setup: "default setup",
	}
}

func (t *TestOptions) Set(data interface{}) {
	d := data.(*TestOptions)
	t.Setup = d.Setup
}

var FileOptions *TestOptions = &TestOptions{
	Setup: "file setup",
}

var FileYml string = "mockup/test_options.yml"

func Yml() (string, error) {
	err := options.ToYmlFile(FileOptions, FileYml)
	if err != nil {
		return "", err
	}

	return FileYml, nil
}

var EnvPrefix string = "ok"
var EnvSetupValue = "env setup"
var EnvKey string = fmt.Sprintf("%s%sSETUP", strings.ToUpper(EnvPrefix), options.EnvSeparator)

type TestBadOptions struct {
	Setup func()
}

func (t *TestBadOptions) Default() interface{} {
	return &TestBadOptions{
		Setup: func() {},
	}
}

func (t *TestBadOptions) Set(data interface{}) {
	d := data.(*TestBadOptions)
	t.Setup = d.Setup
}
