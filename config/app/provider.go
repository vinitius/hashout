package app

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewAppConfig,
	wire.Bind(new(Config), new(HashoutApp)),
)
