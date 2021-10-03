package checkout

type Cart struct {
	Items                   []Item
	TotalAmount             uint32
	TotalDiscount           uint32
	TotalAmountWithDiscount uint32
}

func (c *Cart) AddGift(p Product) {
	if yes, _ := c.ContainsGift(); !yes {
		item := Item{Quantity: 1}
		gift := item.Merge(c.GiftFrom(p))
		c.Items = append(c.Items, gift)
	}

}

func (c Cart) GiftFrom(p Product) Product {
	return Product{
		ID:     p.ID,
		Title:  p.Title,
		IsGift: p.IsGift,
		Amount: 0,
	}
}

func (c Cart) ContainsGift() (contains bool, count int) {
	for _, i := range c.Items {
		if i.Product.IsGift {
			contains = true
			count += int(i.Quantity)
		}
	}
	return
}

func (c *Cart) CalculateTotals() {
	c.TotalAmount = 0
	c.TotalAmountWithDiscount = 0
	c.TotalDiscount = 0

	for _, i := range c.Items {
		c.TotalAmount += i.TotalAmount
		c.TotalDiscount += i.Discount
	}

	c.TotalAmountWithDiscount = c.TotalAmount - c.TotalDiscount
}
