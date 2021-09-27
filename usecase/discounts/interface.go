package discounts

import "viniti.us/hashout/models/checkout"

type Reader interface {
	GetDiscount(item *checkout.Item) error
}

type Client interface {
	Reader
}

type Service interface {
	CalculateDiscounts(items []checkout.Item) (itemsWithDiscount []checkout.Item, err error)
}
