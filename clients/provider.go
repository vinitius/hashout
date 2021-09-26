package clients

import (
	"viniti.us/hashout/usecase/discounts"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewDiscountClient,
	wire.Bind(new(discounts.Client), new(DiscountClient)),
)
