package db

import "github.com/google/wire"

var Provider = wire.NewSet(
	NewConnection,
	wire.Bind(new(Database), new(ProductDataset)),
)
