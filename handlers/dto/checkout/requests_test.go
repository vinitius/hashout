package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"viniti.us/hashout/models/checkout"
)

type RequestsTestSuite struct {
	suite.Suite
}

func TestRequestsRun(t *testing.T) {
	s := new(RequestsTestSuite)
	suite.Run(t, s)
}

func (s *RequestsTestSuite) TestItemToDomain() {
	dto := Item{
		ID:       1,
		Quantity: 1,
	}

	expected := checkout.Item{
		Product:  checkout.Product{ID: 1},
		Quantity: 1,
	}

	actual := dto.ToDomain()

	assert.Equal(s.T(), expected, actual)
}

func (s *RequestsTestSuite) TestCheckoutToDomain() {
	i1 := Item{
		ID:       1,
		Quantity: 1,
	}

	i2 := Item{
		ID:       3,
		Quantity: 10,
	}

	items := []Item{i1, i2}
	dto := Checkout{
		Items: items,
	}

	expectedItem1 := checkout.Item{
		Product:  checkout.Product{ID: 1},
		Quantity: 1,
	}

	expectedItem2 := checkout.Item{
		Product:  checkout.Product{ID: 3},
		Quantity: 10,
	}

	expectedItems := []checkout.Item{expectedItem1, expectedItem2}

	expected := checkout.Cart{
		TotalAmount:             0,
		TotalDiscount:           0,
		TotalAmountWithDiscount: 0,
		Items:                   expectedItems,
	}

	actual := dto.ToDomain()

	assert.Equal(s.T(), expected, actual)
}

func (s *RequestsTestSuite) TestCheckoutWithRepeatedItemsToDomain() {
	i1 := Item{
		ID:       1,
		Quantity: 2,
	}

	i2 := Item{
		ID:       1,
		Quantity: 10,
	}

	items := []Item{i1, i2}
	dto := Checkout{
		Items: items,
	}

	expectedItem1 := checkout.Item{
		Product:  checkout.Product{ID: 1},
		Quantity: 12,
	}

	expectedItems := []checkout.Item{expectedItem1}

	expected := checkout.Cart{
		TotalAmount:             0,
		TotalDiscount:           0,
		TotalAmountWithDiscount: 0,
		Items:                   expectedItems,
	}

	actual := dto.ToDomain()

	assert.Equal(s.T(), expected, actual)
}
