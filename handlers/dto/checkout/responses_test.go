package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"viniti.us/hashout/test/factories"
)

type ResponsesTestSuite struct {
	suite.Suite
}

func TestResponsesRun(t *testing.T) {
	s := new(ResponsesTestSuite)
	suite.Run(t, s)
}

func (s *ResponsesTestSuite) TestToItemResponse() {
	domain := factories.NewItem()

	expectedResponse := ItemResponse{
		ID:          domain.Product.ID,
		Quantity:    domain.Quantity,
		UnitAmount:  domain.UnitAmount,
		TotalAmount: domain.TotalAmount,
		Discount:    domain.Discount,
		IsGift:      domain.Product.IsGift,
	}

	actual := ToItemResponse(*domain)

	assert.Equal(s.T(), expectedResponse, actual)
}

func (s *ResponsesTestSuite) TestToCheckoutResponse() {
	domain := factories.NewCart()

	var items []ItemResponse
	for _, i := range domain.Items {
		items = append(items, ToItemResponse(i))
	}
	expectedResponse := CheckoutResponse{
		TotalAmount:             domain.TotalAmount,
		TotalDiscount:           domain.TotalDiscount,
		TotalAmountWithDiscount: domain.TotalAmountWithDiscount,
		Items:                   items,
	}

	actual := ToCheckoutResponse(*domain)

	assert.Equal(s.T(), expectedResponse, actual)
}
