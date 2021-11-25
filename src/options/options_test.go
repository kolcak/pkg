package options_test

import (
	"os"
	"testing"

	"github.com/kolcak/pkg/src/options"
	mockup_test "github.com/kolcak/pkg/src/options/mockup"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	template := new(mockup_test.TestOptions)
	opt := options.New()

	err := opt.Load(template)
	assert.Nil(t, err)
	assert.Equal(t, template.Default(), template)
}

func TestWithFile(t *testing.T) {
	file, err := mockup_test.Yml()
	assert.Nil(t, err)
	defer os.Remove(file)

	template := new(mockup_test.TestOptions)
	opt := options.New()
	opt.WithFile(file)

	err = opt.Load(template)
	assert.Nil(t, err)
	assert.Equal(t, mockup_test.FileOptions.Setup, template.Setup)
}

func TestWithEnv(t *testing.T) {
	os.Setenv(mockup_test.EnvKey, mockup_test.EnvSetupValue)
	defer os.Unsetenv(mockup_test.EnvKey)

	template := new(mockup_test.TestOptions)
	opt := options.New()
	opt.WithEnv(mockup_test.EnvPrefix)

	err := opt.Load(template)
	assert.Nil(t, err)

	assert.Equal(t, mockup_test.EnvSetupValue, template.Setup)
}

func TestWithBoth(t *testing.T) {
	os.Setenv(mockup_test.EnvKey, mockup_test.EnvSetupValue)
	defer os.Unsetenv(mockup_test.EnvKey)

	file, err := mockup_test.Yml()
	assert.Nil(t, err)
	defer os.Remove(file)

	template := new(mockup_test.TestOptions)
	opt := options.New()
	opt.
		WithEnv(mockup_test.EnvPrefix).
		WithFile(file)

	err = opt.Load(template)
	assert.Nil(t, err)

	assert.NotEqual(t, mockup_test.FileOptions.Setup, template.Setup)
	assert.Equal(t, mockup_test.EnvSetupValue, template.Setup)
}

func TestWithErrorDefaultJsonMarshal(t *testing.T) {
	template := new(mockup_test.TestBadOptions)
	opt := options.New()

	err := opt.Load(template)
	assert.NotNil(t, err)
}

func TestWithErrorReadInConfig(t *testing.T) {
	template := new(mockup_test.TestOptions)
	opt := options.New()

	opt.WithFile("notExist.yml")
	err := opt.Load(template)
	assert.NotNil(t, err)
}
