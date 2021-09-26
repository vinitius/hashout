package handlers

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewCheckoutHandler,
)
