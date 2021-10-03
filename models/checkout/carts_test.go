package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CartsTestSuite struct {
	suite.Suite
}

func TestCartsRun(t *testing.T) {
	s := new(CartsTestSuite)
	suite.Run(t, s)
}

func (s *CartsTestSuite) TestCalculateTotals() {
	p1 := Product{ID: 1, Amount: 2000, IsGift: false}
	i1 := Item{
		Product:      p1,
		Quantity:     2,
		DiscountRate: 0.10,
		UnitAmount:   p1.Amount,
		TotalAmount:  p1.Amount * 2,
	}
	i1.CalculateDiscount()
	p2 := Product{ID: 2, Amount: 3000, IsGift: false}
	i2 := Item{
		Product:      p2,
		Quantity:     3,
		DiscountRate: 0.10,
		UnitAmount:   p2.Amount,
		TotalAmount:  p2.Amount * 3,
	}
	i2.CalculateDiscount()
	items := []Item{i1, i2}
	cart := Cart{
		Items: items,
	}

	expectedCart := Cart{
		Items:                   items,
		TotalAmount:             i1.TotalAmount + i2.TotalAmount,
		TotalDiscount:           i1.Discount + i2.Discount,
		TotalAmountWithDiscount: (i1.TotalAmount + i2.TotalAmount) - (i1.Discount + i2.Discount),
	}

	cart.CalculateTotals()

	assert.Exactly(s.T(), expectedCart, cart)
}
