package clients

import (
	"context"

	"viniti.us/hashout/config/cb"
	"viniti.us/hashout/config/log"

	"viniti.us/hashout/models/checkout"
	discount "viniti.us/hashout/pb"
)

type DiscountClient struct {
	grpc    discount.DiscountClient
	ctx     context.Context
	command cb.DiscountCB
}

func NewDiscountClient(ctx context.Context, grpc discount.DiscountClient, command cb.DiscountCB) DiscountClient {
	return DiscountClient{ctx: ctx, grpc: grpc, command: command}
}

func (c DiscountClient) GetDiscount(item *checkout.Item) error {
	req := &discount.GetDiscountRequest{
		ProductID: item.Product.ID,
	}

	res, err := c.doGetDiscount(req)
	if err != nil {
		log.Logger.Errorf("error getting discount for product %d", item.Product.ID, err)
		return err
	}

	log.Logger.Infof("Discount rate for product %d is %.2f -> %.2f%% ", item.Product.ID, res.GetPercentage(), res.GetPercentage()*100)

	item.DiscountRate = res.GetPercentage() * 100 // notation convenience

	return nil
}

func (c DiscountClient) doGetDiscount(req *discount.GetDiscountRequest) (*discount.GetDiscountResponse, error) {

	d, err := c.command.CB.Do(c.ctx, func() (interface{}, error) {
		return c.grpc.GetDiscount(c.ctx, req)
	})

	if err != nil {
		return nil, err
	}

	res := d.(*discount.GetDiscountResponse)

	return res, err
}
