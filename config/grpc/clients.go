package grpc

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	discount "viniti.us/hashout/pb"
)

func NewDiscountGRPCClient() discount.DiscountClient {
	conn, err := grpc.Dial(viper.GetString("DISCOUNT_GRPC_SERVER_ADDRESS"), grpc.WithInsecure()) // set TLS conf accordingly
	if err != nil {
		panic(err)
	}

	client := discount.NewDiscountClient(conn)
	return client
}

func NewContext() context.Context {
	return context.Background()
}
