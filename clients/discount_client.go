package clients

import (
	"context"
)

type DiscountClient struct {
	ctx context.Context
}

func NewDiscountClient(ctx context.Context) DiscountClient {
	return DiscountClient{ctx: ctx}
}
