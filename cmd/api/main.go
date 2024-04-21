package main

import (
	"os"

	"github.com/appstore-notify-sample/pkg/cmd/api"
)

func main() {
	os.Exit(run())
}

func run() int {
	api.Run()
	return 0
}
