package cb

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewDiscountCB,
)
