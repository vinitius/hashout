package discounts

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"viniti.us/hashout/config/log"
	domain "viniti.us/hashout/models/checkout"
	customErr "viniti.us/hashout/models/errors"
	"viniti.us/hashout/test/factories"
	mocks "viniti.us/hashout/test/mocks/clients"
)

const (
	getDicountMethod = "GetDiscount"
)

type DiscountsUseCaseSuite struct {
	suite.Suite
	cli     *mocks.DiscountClient
	useCase UseCase
}

func (s *DiscountsUseCaseSuite) SetupTest() {
	log.SetupLogger()
	s.cli = new(mocks.DiscountClient)
	s.useCase = UseCase{cli: s.cli}
}

func TestRun(t *testing.T) {
	s := new(DiscountsUseCaseSuite)
	suite.Run(t, s)
}

func (s *DiscountsUseCaseSuite) TestCalculateDiscountsSuccessfully() {
	item1 := factories.NewItem()
	item2 := factories.NewItem()
	items := []domain.Item{*item1, *item2}

	expectedItem1 := domain.Item{
		Product:      item1.Product,
		Quantity:     item1.Quantity,
		DiscountRate: item1.DiscountRate,
		UnitAmount:   item1.UnitAmount,
		TotalAmount:  item1.TotalAmount,
	}
	expectedItem1.CalculateDiscount()
	expectedItem2 := domain.Item{
		Product:      item2.Product,
		Quantity:     item2.Quantity,
		DiscountRate: item2.DiscountRate,
		UnitAmount:   item2.UnitAmount,
		TotalAmount:  item2.TotalAmount,
	}
	expectedItem2.CalculateDiscount()
	expectedItems := []domain.Item{expectedItem1, expectedItem2}

	s.cli.On(getDicountMethod, item1).Return(nil)
	s.cli.On(getDicountMethod, item2).Return(nil)

	itemsWithDiscounts, err := s.useCase.CalculateDiscounts(items)

	assert.Exactly(s.T(), expectedItems, itemsWithDiscounts)
	assert.Nil(s.T(), err)
	s.cli.AssertNumberOfCalls(s.T(), getDicountMethod, len(items))
}

func (s *DiscountsUseCaseSuite) TestCalculateDiscountsError() {
	item1 := factories.NewItem()
	item2 := factories.NewItem()
	items := []domain.Item{*item1, *item2}

	expectedItem1 := domain.Item{
		Product:      item1.Product,
		Quantity:     item1.Quantity,
		DiscountRate: item1.DiscountRate,
		UnitAmount:   item1.UnitAmount,
		TotalAmount:  item1.TotalAmount,
	}
	expectedItem1.CalculateDiscount()
	expectedItem2 := domain.Item{
		Product:      item2.Product,
		Quantity:     item2.Quantity,
		DiscountRate: item2.DiscountRate,
		UnitAmount:   item2.UnitAmount,
		TotalAmount:  item2.TotalAmount,
	}
	expectedItems := []domain.Item{expectedItem1, expectedItem2}
	expectedError := errors.New("could not get discounts for products [" + strconv.Itoa(int(item2.Product.ID)) + "]")
	finalError := &customErr.DiscountError{Type: "Fetch", Err: expectedError}

	s.cli.On(getDicountMethod, item1).Return(nil)
	s.cli.On(getDicountMethod, item2).Return(expectedError)

	itemsWithDiscounts, err := s.useCase.CalculateDiscounts(items)

	assert.Exactly(s.T(), expectedItems, itemsWithDiscounts)
	assert.Exactly(s.T(), finalError, err)
	s.cli.AssertNumberOfCalls(s.T(), getDicountMethod, len(items))
}
