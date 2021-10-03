package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductsTestSuite struct {
	suite.Suite
}

func TestProductsRun(t *testing.T) {
	s := new(ProductsTestSuite)
	suite.Run(t, s)
}

func (s *ProductsTestSuite) TestMerge() {
	p1 := Product{ID: 1, Amount: 2000, IsGift: false}
	item := Item{
		Product:  Product{ID: 1},
		Quantity: 2,
	}
	expected := Item{
		Product:      p1,
		Quantity:     2,
		DiscountRate: 0,
		UnitAmount:   p1.Amount,
		TotalAmount:  p1.Amount * 2,
	}

	assert.Exactly(s.T(), expected, item.Merge(p1))
}

func (s *ProductsTestSuite) TestCalculateDiscount() {
	p1 := Product{ID: 1, Amount: 2000, IsGift: false}
	item := Item{
		Quantity:     2,
		DiscountRate: 10,
	}
	item = item.Merge(p1)

	expected := Item{
		Product:      p1,
		Quantity:     2,
		DiscountRate: 10,
		UnitAmount:   p1.Amount,
		TotalAmount:  p1.Amount * 2,
		Discount:     400, //10% rate on 2000(unit)*2(qty)
	}

	item.CalculateDiscount()

	assert.Exactly(s.T(), expected, item)
}
