# Prom init

Simple library to automate prometheus metrics serving.

Example:

```go
package main

import (
	_ "github.com/Dalee/prom-init"
	"time"
)

func main() {
	time.Sleep(time.Second * 100)
}
```

Now you have a running prometheus on `:7070`.
