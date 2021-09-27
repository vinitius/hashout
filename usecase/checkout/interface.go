package checkout

import "viniti.us/hashout/models/checkout"

type Reader interface {
	Find(items []checkout.Item) (mergedItems []checkout.Item, err error)
}

type Repository interface {
	Reader
}

type Service interface {
	Checkout(c *checkout.Cart) error
}
