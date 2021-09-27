package dto

import "viniti.us/hashout/models/checkout"

type Item struct {
	ID       int32  `json:"id" binding:"required,gt=0"`
	Quantity uint32 `json:"quantity" binding:"required,gt=0"`
}

type Checkout struct {
	Items []Item `json:"products" binding:"required,dive"`
}

func (i Item) ToDomain() checkout.Item {
	return checkout.Item{
		Product:  checkout.Product{ID: i.ID},
		Quantity: i.Quantity,
	}
}

func (c Checkout) ToDomain() checkout.Cart {
	var items []checkout.Item
	uniqueItems := make(map[int32]Item)
	for _, i := range c.Items {
		if repeated, found := uniqueItems[i.ID]; found {
			uniqueItems[i.ID] = Item{ID: repeated.ID, Quantity: repeated.Quantity + i.Quantity}
		} else {
			uniqueItems[i.ID] = i
		}
	}

	for _, i := range uniqueItems {
		items = append(items, i.ToDomain())
	}

	return checkout.Cart{
		Items: items,
	}
}
