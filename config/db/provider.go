package db

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewConnection,
)
