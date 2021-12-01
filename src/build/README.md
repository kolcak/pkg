#### Install

```go
go get github.com/kolcak/pkg/src/build
```

#### Use

```bash
go run -ldflags "-w -s -X build.xProject=project -X build.xVersion=`date +"v%Y.%m%d.%H%M%S"` -X build.xRevision=`date +"v%Y.%m%d.%H%M%S"` -X build.xRelease=`date -u +%s` -X build.xEnv=IWM" main.go
```

```go
package main

import (
	"fmt"

	"github.com/kolcak/pkg/src/build"
)

func main() {
	fmt.Printf("%+v\n", build.Info())
}
```
