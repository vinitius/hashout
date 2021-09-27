package checkout

type Product struct {
	ID          int32
	Title       string
	Description string
	Amount      uint32
	IsGift      bool `json:"is_gift"`
}

type Item struct {
	Product      Product
	Quantity     uint32
	DiscountRate uint32
	UnitAmount   uint32
	TotalAmount  uint32
	Discount     uint32
}

func (i Item) Merge(p Product) Item {
	return Item{
		Product:      p,
		Quantity:     i.Quantity,
		DiscountRate: i.DiscountRate,
		UnitAmount:   p.Amount,
		TotalAmount:  p.Amount * i.Quantity,
		Discount:     i.Discount,
	}
}
