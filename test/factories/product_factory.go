package factories

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	domain "viniti.us/hashout/models/checkout"
)

func NewItem() *domain.Item {
	rand.Seed(time.Now().UnixNano())
	qty := rand.Intn(20-1) + 1
	p := NewProduct()
	return &domain.Item{
		Product:      *p,
		Quantity:     uint32(qty),
		DiscountRate: 5,
		TotalAmount:  p.Amount * uint32(qty),
		UnitAmount:   p.Amount,
	}

}

func NewProduct() *domain.Product {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(6-1) + 1
	price := rand.Intn(5000-500) + 500
	return &domain.Product{
		ID:          int32(id),
		Title:       uuid.New().String(),
		Amount:      uint32(price),
		Description: uuid.New().String(),
		IsGift:      false,
	}

}
