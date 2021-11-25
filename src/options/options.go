package options

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const EnvSeparator = "_"

type Options interface {
	Default() interface{}
	Set(data interface{})
}

type O struct {
	template interface{}
	reader   *viper.Viper
	loaded   bool
	file     string
	prefix   string
}

func New() *O {
	return &O{
		reader: viper.New(),
	}
}

func (o *O) WithFile(file string) *O {
	o.file = file
	return o
}

func (o *O) WithEnv(prefix string) *O {
	o.prefix = strings.ToUpper(prefix)
	return o
}

func (o *O) withDefault(template interface{}) error {
	Log.Debugf("withDefault: %+v", template)
	if cfgJson, err := json.Marshal(template); err != nil {
		return fmt.Errorf("withDefault: Marshal default options, error: %+v", err)
	} else {
		o.reader.SetConfigType("json")
		if err := o.reader.ReadConfig(bytes.NewBuffer(cfgJson)); err != nil {
			return fmt.Errorf("withDefault: Load default options, error: %+v", err)
		} else {
			if err := o.reader.Unmarshal(template); err != nil {
				return fmt.Errorf("withDefault: Unable to decode into struct, error: %+v", err)
			}
		}
	}
	return nil
}

func (o *O) withFile(template interface{}) error {
	if o.file != "" {
		Log.Debugf("withFile: `%s`", o.file)
		ext := filepath.Ext(o.file)
		if len(ext) > 1 {
			o.reader.SetConfigType(ext[1:])
		}
		o.reader.SetConfigFile(o.file)
		if err := o.reader.ReadInConfig(); err != nil {
			return err
		} else {
			if err := o.reader.Unmarshal(template); err != nil {
				return fmt.Errorf("withFile: Unable to decode into struct, error: %+v", err)
			}
		}
	}
	return nil
}

func (o *O) withEnv() {
	if o.prefix != "" {
		Log.Debugf("with env prefix: %s", o.prefix)
		o.reader.SetEnvPrefix(o.prefix)
		o.reader.SetEnvKeyReplacer(strings.NewReplacer(".", EnvSeparator))
		o.reader.AutomaticEnv()
	}
}

func (o *O) Load(template Options) error {
	tmp := template.Default()
	o.withEnv()
	if err := o.withDefault(&tmp); err != nil {
		return err
	}

	if err := o.withFile(&tmp); err != nil {
		return err
	}

	template.Set(tmp)
	return nil
}

// Logger
type Logger interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type l struct{}

func (l) Infof(format string, args ...interface{})  {}
func (l) Debugf(format string, args ...interface{}) {}

var Log Logger = new(l)
