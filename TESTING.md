# How to test
For testing things that require debugging mode, or a running app inside the project folder a **test.go** file can be used, such as:
``` Go
package main

import (
	"fmt"
	"io"
	"time"

	"github.com/Lunal98/netwatchdog/cmd/check"
)

func main() {


	var nwd NwdCore
	Interface := check.InterfaceCheck{
		CheckName: "eth0",
	}
	nwd.AddCheck(&Interface, time.Minute, "Interface Check", 1)
	nwd.Start()

}
```