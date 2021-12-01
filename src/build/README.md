#### Install

```go
go get github.com/kolcak/pkg/src/build
```

#### Use

```bash
# OR BUILD
go run -ldflags "-w -s -X github.com/kolcak/pkg/src/build.xProject=project -X github.com/kolcak/pkg/src/build.xVersion=`date +"v%Y.%m%d.%H%M%S"` -X github.com/kolcak/pkg/src/build.xRevision=`date +"v%Y.%m%d.%H%M%S"` -X github.com/kolcak/pkg/src/build.xRelease=`date -u +%s` -X github.com/kolcak/pkg/src/build.xEnv=IWM" main.go
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
