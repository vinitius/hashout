package discounts

import (
	"encoding/json"
	"errors"

	"viniti.us/hashout/config/log"
	"viniti.us/hashout/models/checkout"
	customErr "viniti.us/hashout/models/errors"
)

type UseCase struct {
	cli Client
}

func NewUseCase(cli Client) UseCase {
	return UseCase{cli: cli}
}

func (u UseCase) CalculateDiscounts(items []checkout.Item) (itemsWithDiscount []checkout.Item, err error) {
	var failedDiscounts []int32
	for _, i := range items {
		if err := u.cli.GetDiscount(&i); err != nil {
			log.Logger.Warnf("could not get discount for product %d", i.Product.ID, err)
			failedDiscounts = append(failedDiscounts, i.Product.ID)
		} else {
			i.CalculateDiscount()
			itemsWithDiscount = append(itemsWithDiscount, i)
		}
	}

	if len(failedDiscounts) > 0 {
		p, _ := json.Marshal(failedDiscounts)
		err = &customErr.DiscountError{Type: "Fetch", Err: errors.New("could not get discounts for products " + string(p))}
	}

	return
}
