#### Install

```go
go get github.com/kolcak/pkg/src/options
```

#### Use

```go
package main

import (
	"fmt"

	"github.com/kolcak/pkg/src/options"
)

type Options struct {
	Setup string
}

func (*Options) Default() interface{} {
	return &Options{
		Setup: "default",
	}
}

func (o *Options) Set(data interface{}) {
	d := data.(*Options)
	o.Setup = d.Setup
}

func main() {
	template := new(Options)
	opt := options.New()
	opt.
		WithEnv("ts"). // apply env. variables: TS_*
		WithFile("config.yml").
		Load(template)

	fmt.Printf("%+v\n", template.Default())
	fmt.Printf("%+v\n", template)
}
```
