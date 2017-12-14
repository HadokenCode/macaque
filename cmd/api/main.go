package main

import (
	"fmt"

	"github.com/wildnature/macaque/pkg/server/api"
)

func main() {
	fmt.Println("Launching server...")
	api.ConfigureServerAndRun()
}
