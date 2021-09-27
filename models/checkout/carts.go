package checkout

type Cart struct {
	Items                   []Item
	TotalAmount             uint32
	TotalDiscount           uint32
	TotalAmountWithDiscount uint32
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
