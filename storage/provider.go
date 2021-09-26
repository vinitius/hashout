package storage

import (
	"viniti.us/hashout/usecase/checkout"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewProductRepository,
	wire.Bind(new(checkout.Repository), new(ProductRepository)),
)
