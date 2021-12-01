package options

import "os"

func GetEnv(key string, fn func(v string) (r interface{})) interface{} {
	return fn(os.Getenv(key))
}
