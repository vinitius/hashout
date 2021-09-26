package usecase

import (
	"viniti.us/hashout/usecase/checkout"
	"viniti.us/hashout/usecase/discounts"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	checkout.NewUseCase,
	wire.Bind(new(checkout.Service), new(checkout.UseCase)),
	discounts.NewUseCase,
	wire.Bind(new(discounts.Service), new(discounts.UseCase)),
)
