package main

import (
  "github.com/cli-server/route"
)

func main() {
  e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
