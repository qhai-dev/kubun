package transport

import (
	"github.com/google/wire"
	"github.com/qhai-dev/ozma/foundation/transport/rpc"
)

var Provider = wire.NewSet(
	rpc.NewUserServer,
	rpc.NewGRPCServer,
)
