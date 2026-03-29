package main

import (
	"os"

	"github.com/qhai-dev/ozma/cmd/foundation/app"

	"github.com/qhai-dev/ozma/pkg/cli"
)

func main() {
	command := app.NewFoundationServerCommand()
	code := cli.Run(command)
	os.Exit(code)
}
