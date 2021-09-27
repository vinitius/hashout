package dto

import "viniti.us/hashout/models/checkout"

type CheckoutResponse struct {
	Items                   []ItemResponse `json:"products"`
	TotalAmount             uint32         `json:"total_amount"`
	TotalDiscount           uint32         `json:"total_discount"`
	TotalAmountWithDiscount uint32         `json:"total_amount_with_discount"`
}

type ItemResponse struct {
	ID          int32  `json:"id"`
	Quantity    uint32 `json:"quantity"`
	UnitAmount  uint32 `json:"unit_amount"`
	TotalAmount uint32 `json:"total_amount"`
	Discount    uint32 `json:"discount"`
	IsGift      bool   `json:"is_gift"`
}

func ToItemResponse(i checkout.Item) ItemResponse {
	return ItemResponse{
		ID:          i.Product.ID,
		Quantity:    i.Quantity,
		UnitAmount:  i.UnitAmount,
		TotalAmount: i.TotalAmount,
		Discount:    i.Discount,
		IsGift:      i.Product.IsGift,
	}
}

func ToCheckoutResponse(c checkout.Cart) CheckoutResponse {
	var items []ItemResponse
	for _, i := range c.Items {
		items = append(items, ToItemResponse(i))
	}

	return CheckoutResponse{
		TotalAmount:             c.TotalAmount,
		TotalDiscount:           c.TotalDiscount,
		TotalAmountWithDiscount: c.TotalAmountWithDiscount,
		Items:                   items,
	}
}
