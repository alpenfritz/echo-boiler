package main

import (
	"echo-boiler/router"
)

func main() {
	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
