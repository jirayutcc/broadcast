package main

import (
	"github.com/jirayutcc/transaction-broadcast/core"
)

func main() {
	e := core.Setup()

	e.Logger.Fatal(e.Start(":3033"))
}
