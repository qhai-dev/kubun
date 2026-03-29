//go:build wireinject
// +build wireinject

package foundation

import (
	"github.com/google/wire"

	"github.com/qhai-dev/ozma/foundation/application"
	"github.com/qhai-dev/ozma/foundation/infra"
	"github.com/qhai-dev/ozma/foundation/transport"
)

func InitializeApp() (*App, error) {
	wire.Build(
		infra.Provider,
		application.Provider,
		transport.Provider,
		NewApp,
	)

	return  &App{}, nil
}
