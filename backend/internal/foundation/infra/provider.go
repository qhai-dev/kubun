package infra

import (
	"github.com/google/wire"
	"github.com/qhai-dev/ozma/foundation/infra/database"
	"github.com/qhai-dev/ozma/foundation/infra/opentelemetryx"
	"github.com/qhai-dev/ozma/foundation/infra/repository"
)

var Provider = wire.NewSet(
	database.NewDatabase,
	repository.NewUserRepository,
	opentelemetryx.NewOpenTelemetry,
)
