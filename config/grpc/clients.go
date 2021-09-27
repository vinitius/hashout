package grpc

import (
	"context"

	discount "viniti.us/hashout/pb"
)

func NewDiscountGRPCClient() discount.DiscountClient {
	return nil
}

func NewContext() context.Context {
	return context.Background()
}
