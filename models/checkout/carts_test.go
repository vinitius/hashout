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

func (s *CartsTestSuite) TestContainsGift() {
	gift := Product{ID: 1, Amount: 2000, IsGift: true}
	i1 := Item{
		Product:      gift,
		Quantity:     2,
		DiscountRate: 0.10,
		UnitAmount:   gift.Amount,
		TotalAmount:  gift.Amount * 2,
	}
	i1 = i1.Merge(gift)
	p2 := Product{ID: 2, Amount: 3000, IsGift: false}
	i2 := Item{
		Product:      p2,
		Quantity:     3,
		DiscountRate: 0.10,
		UnitAmount:   p2.Amount,
		TotalAmount:  p2.Amount * 3,
	}
	i2 = i2.Merge(p2)
	cartWithGift := Cart{
		Items: []Item{i1, i2},
	}
	cartWithoutGift := Cart{
		Items: []Item{i2},
	}

	tests := []struct {
		name   string
		qty    int
		cart   Cart
		wanted bool
	}{
		{
			name:   "with gift",
			cart:   cartWithGift,
			qty:    2,
			wanted: true,
		},
		{
			name:   "without gift",
			cart:   cartWithoutGift,
			qty:    0,
			wanted: false,
		},
	}

	for _, c := range tests {
		s.T().Run(c.name, func(t *testing.T) {
			contains, count := c.cart.ContainsGift()
			assert.Equal(s.T(), c.wanted, contains)
			assert.Equal(s.T(), c.qty, count)
		})
	}

}

func (s *CartsTestSuite) TestGiftFrom() {
	gift := Product{ID: 1, Amount: 2000, IsGift: true}

	expectedProduct := Product{
		ID:          gift.ID,
		Title:       gift.Title,
		Description: gift.Description,
		Amount:      0,
		IsGift:      gift.IsGift,
	}

	cart := Cart{}

	assert.Exactly(s.T(), expectedProduct, cart.GiftFrom(gift))
}

func (s *CartsTestSuite) TestAddGift() {
	gift := Product{ID: 1, Amount: 2000, IsGift: true}
	i1 := Item{
		Product:      gift,
		Quantity:     2,
		DiscountRate: 0.10,
		UnitAmount:   gift.Amount,
		TotalAmount:  gift.Amount * 2,
	}
	i1 = i1.Merge(gift)
	p2 := Product{ID: 2, Amount: 3000, IsGift: false}
	i2 := Item{
		Product:      p2,
		Quantity:     3,
		DiscountRate: 0.10,
		UnitAmount:   p2.Amount,
		TotalAmount:  p2.Amount * 3,
	}
	i2 = i2.Merge(p2)
	cartWithGift := Cart{
		Items: []Item{i1, i2},
	}
	cartWithoutGift := Cart{
		Items: []Item{i2},
	}

	expectedGiftItem := Item{
		Product:      cartWithGift.GiftFrom(gift),
		Quantity:     1,
		UnitAmount:   0,
		TotalAmount:  0,
		DiscountRate: 0,
		Discount:     0,
	}
	expectedWithExtraItem := Cart{
		Items: []Item{i2, expectedGiftItem},
	}

	tests := []struct {
		name   string
		cart   Cart
		gift   Product
		wanted Cart
	}{
		{
			name:   "with gift",
			cart:   cartWithGift,
			gift:   gift,
			wanted: cartWithGift,
		},
		{
			name:   "without gift",
			cart:   cartWithoutGift,
			gift:   gift,
			wanted: expectedWithExtraItem,
		},
	}

	for _, c := range tests {
		s.T().Run(c.name, func(t *testing.T) {
			c.cart.AddGift(c.gift)
			assert.Exactly(s.T(), c.wanted, c.cart)
		})
	}

}
