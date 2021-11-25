package options

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

func ToYmlFile(o Options, fileName string) error {
	Log.Infof("Yml.ToFile: Start generate file %s", fileName)
	if es, err := yaml.Marshal(o); err != nil {
		return err
	} else {
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer func() error {
			if err := f.Close(); err != nil {
				return fmt.Errorf("Yml.ToFile: Close file error: %+v", err)
			}
			return nil
		}()

		fmt.Fprintln(f, "# This file is auto-generated from config")
		fmt.Fprint(f, string(es))

		Log.Infof("Yml.ToFile: done! %s", fileName)
	}

	return nil
}

func ToEnvFile(o Options, prefix string, fileName string) error {
	Log.Infof("Env.ToFile: Start generate file %s", fileName)
	if es, err := marshaler(o.Default(), prefix, 0, make([]string, 0), make([]string, 0)); err != nil {
		return err
	} else {
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}

		defer func() error {
			if err := f.Close(); err != nil {
				return fmt.Errorf("ToEnvFile: Close file error: %+v", err)
			}

			return nil
		}()

		fmt.Fprintln(f, "# This file is auto-generated from config")
		fmt.Fprintln(f, "# if you need to modify a variable, create: .env")

		for _, value := range es {
			fmt.Fprint(f, value)
		}
		Log.Infof("ToEnvFile: prefix '%s' file: '%s' done!", prefix, fileName)
	}

	return nil
}

func marshaler(v interface{}, prefix string, lvl int, expose []string, ignore []string) ([]string, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("env.marshaler: value must be a non-nil pointer to a struct")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("env.marshaler: value must be a non-nil pointer to a struct")
	}

	output := []string{}
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		valueField := rv.Field(i)
		typeField := t.Field(i)
		nameField := strings.ToUpper(fmt.Sprintf("%s%s%s", prefix, EnvSeparator, typeField.Name))
		tag := typeField.Tag.Get("envDesc")

		if tag == "-" {
			continue
		}

		if len(expose) > 0 {
			skip := true
			for _, val := range expose {
				if val == typeField.Name {
					skip = false
				}
			}

			if skip {
				continue
			}
		}

		if len(ignore) > 0 {
			skip := false
			for _, val := range ignore {
				if val == typeField.Name {
					skip = true
				}
			}

			if skip {
				continue
			}
		}

		tag = strings.ReplaceAll(tag, "\n", "\n##")

		switch valueField.Kind() {
		case reflect.Struct:
			if !valueField.Addr().CanInterface() {
				continue
			}

			iface := valueField.Addr().Interface()
			ex := make([]string, 0)
			if typeField.Tag.Get("envExpose") != "" {
				ex = strings.Split(typeField.Tag.Get("envExpose"), ",")
			}

			ig := make([]string, 0)
			if typeField.Tag.Get("envIgnore") != "" {
				ig = strings.Split(typeField.Tag.Get("envIgnore"), ",")
			}

			if nes, err := marshaler(iface, nameField, 1, ex, ig); err != nil {
				return nil, err
			} else {
				if len(nes) != 0 {
					if lvl == 0 {
						output = append(output, "\n")
					}

					sectionSeparator := strings.Repeat("=", 150)
					if tag != "" {
						if lvl > 0 {
							output = append(output, fmt.Sprintf("## %s\n", tag))
						} else {
							output = append(output, fmt.Sprintf("\n## %s\n", sectionSeparator))
							output = append(output, fmt.Sprintf("## %s\n", tag))
							output = append(output, fmt.Sprintf("## %s\n\n", sectionSeparator))
						}
					} else {
						if lvl == 0 {
							output = append(output, fmt.Sprintf("\n## %s\n", sectionSeparator))
							output = append(output, fmt.Sprintf("## %s\n", typeField.Name))
							output = append(output, fmt.Sprintf("## %s\n\n", sectionSeparator))
						}
					}

					output = append(output, nes...)
				}
			}

		case reflect.Bool:
			output = append(output, tagFormat(nameField, strings.ToUpper(fmt.Sprintf("%v", valueField)), tag))

		default:
			output = append(output, tagFormat(nameField, fmt.Sprintf("%v", valueField), tag))
		}
	}

	return output, nil
}

func tagFormat(key string, value string, tag string) string {
	if tag == "" {
		return fmt.Sprintf("\n%s=%s\n", key, value)
	}

	return fmt.Sprintf("\n## %s\n%s=%s\n", tag, key, value)
}
