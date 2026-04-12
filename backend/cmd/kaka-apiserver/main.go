package main

import (
	"os"

	"github.com/qhai-dev/kaka/cmd/kaka-apiserver/app"
	"github.com/qhai-dev/kaka/component/command"

	_ "github.com/qhai-dev/kaka/component/log"
)

func main() {
	cmd := app.NewAPIServerCommand()
	code := command.Run(cmd)
	os.Exit(code)
}
