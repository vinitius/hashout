package checkout

import "viniti.us/hashout/models/checkout"

type Reader interface {
	FindAll(items []checkout.Item) (mergedItems []checkout.Item, err error)
	FindLastByIsGift(isGift bool) (product checkout.Product, err error)
}

type Repository interface {
	Reader
}

type Service interface {
	Checkout(c *checkout.Cart) error
}
