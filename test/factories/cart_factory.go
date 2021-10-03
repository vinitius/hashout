package factories

import (
	domain "viniti.us/hashout/models/checkout"
)

func NewCart() *domain.Cart {
	i1 := NewItem()
	i2 := NewItem()
	i3 := NewItem()
	items := []domain.Item{*i1, *i2, *i3}
	return &domain.Cart{
		Items: items,
	}

}
