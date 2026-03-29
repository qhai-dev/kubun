package application

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewCreateUserHandler,
)
