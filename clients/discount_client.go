package clients

import (
	"context"

	"viniti.us/hashout/config/log"

	"viniti.us/hashout/models/checkout"
	discount "viniti.us/hashout/pb"
)

type DiscountClient struct {
	grpc discount.DiscountClient
	ctx  context.Context
}

func NewDiscountClient(ctx context.Context, grpc discount.DiscountClient) DiscountClient {
	return DiscountClient{ctx: ctx, grpc: grpc}
}

func (c DiscountClient) GetDiscount(item *checkout.Item) error {
	req := &discount.GetDiscountRequest{
		ProductID: item.Product.ID,
	}

	res, err := c.grpc.GetDiscount(c.ctx, req)
	if err != nil {
		log.Logger.Errorf("error getting discount for product %d", item.Product.ID, err)
		return err
	}

	log.Logger.Infof("Discount rate for product %d is %.2f -> %.2f%% ", item.Product.ID, res.GetPercentage(), res.GetPercentage()*100)

	item.DiscountRate = res.GetPercentage() * 100

	return nil
}
