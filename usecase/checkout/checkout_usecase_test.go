package checkout

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"viniti.us/hashout/config/log"
	domain "viniti.us/hashout/models/checkout"
	customErr "viniti.us/hashout/models/errors"
	"viniti.us/hashout/test/factories"
	mockConfig "viniti.us/hashout/test/mocks/config"
	mocksRepo "viniti.us/hashout/test/mocks/storage"
	mocksUsecase "viniti.us/hashout/test/mocks/usecase"
)

const (
	findAllMethod                = "FindAll"
	findLastByIsGiftMethod       = "FindLastByIsGift"
	calculateDiscountsMethod     = "CalculateDiscounts"
	isBlackFridayActiveMethod    = "IsBlackFridayGiftActive"
	getAllowedGiftsPerCartMethod = "GetAllowedGiftsPerCart"
	getBlackFridayMonthMethod    = "GetBlackFridayMonth"
	getBlackFridayDayMethod      = "GetBlackFridayDay"
)

type CheckoutUseCaseSuite struct {
	suite.Suite
	productRepo     *mocksRepo.ProductRepository
	discountUsecase *mocksUsecase.DiscountUseCase
	config          *mockConfig.HashoutApp
	useCase         UseCase
}

func (s *CheckoutUseCaseSuite) SetupTest() {
	log.SetupLogger()
	s.productRepo = new(mocksRepo.ProductRepository)
	s.discountUsecase = new(mocksUsecase.DiscountUseCase)
	s.config = new(mockConfig.HashoutApp)
	s.useCase = UseCase{repo: s.productRepo, discountsUseCase: s.discountUsecase, config: s.config}
}

func TestRun(t *testing.T) {
	s := new(CheckoutUseCaseSuite)
	suite.Run(t, s)
}

func (s *CheckoutUseCaseSuite) TestCheckoutSuccessfully() {
	cart := factories.NewCart()
	expectedItem1 := factories.NewItem()
	expectedItem2 := factories.NewItem()
	expectedItems := []domain.Item{*expectedItem1, *expectedItem2}
	expectedCart := &domain.Cart{
		Items: expectedItems,
	}
	expectedCart.CalculateTotals()

	s.productRepo.On(findAllMethod, cart.Items).Return(expectedItems, nil)
	s.discountUsecase.On(calculateDiscountsMethod, expectedItems).Return(expectedItems, nil)
	s.config.On(isBlackFridayActiveMethod).Return(false)

	err := s.useCase.Checkout(cart)

	assert.Exactly(s.T(), expectedCart, cart)
	assert.Nil(s.T(), err)
	s.productRepo.AssertNumberOfCalls(s.T(), findAllMethod, 1)
	s.discountUsecase.AssertNumberOfCalls(s.T(), calculateDiscountsMethod, 1)
}

func (s *CheckoutUseCaseSuite) TestCheckoutOnBlackFridaySuccessfully() {
	cart := factories.NewCart()
	expectedItem1 := factories.NewItem()
	expectedItem2 := factories.NewItem()
	expectedGiftProduct := factories.NewProduct()
	expectedGiftProduct.IsGift = true
	expectedItems := []domain.Item{*expectedItem1, *expectedItem2}
	expectedCart := &domain.Cart{
		Items: expectedItems,
	}
	expectedCart.CalculateTotals()

	s.productRepo.On(findAllMethod, cart.Items).Return(expectedItems, nil)
	s.productRepo.On(findLastByIsGiftMethod, true).Return(*expectedGiftProduct, nil)
	s.discountUsecase.On(calculateDiscountsMethod, expectedItems).Return(expectedItems, nil)
	s.config.On(isBlackFridayActiveMethod).Return(true)
	s.config.On(getBlackFridayDayMethod).Return(time.Now().Day())
	s.config.On(getBlackFridayMonthMethod).Return(time.Now().Month())

	err := s.useCase.Checkout(cart)

	expectedCart.AddGift(*expectedGiftProduct)

	assert.Exactly(s.T(), expectedCart, cart)
	assert.Nil(s.T(), err)
	s.productRepo.AssertNumberOfCalls(s.T(), findAllMethod, 1)
	s.productRepo.AssertNumberOfCalls(s.T(), findLastByIsGiftMethod, 1)
	s.discountUsecase.AssertNumberOfCalls(s.T(), calculateDiscountsMethod, 1)
}

func (s *CheckoutUseCaseSuite) TestCheckoutOnBlackFridayError() {
	cart := factories.NewCart()
	expectedItem1 := factories.NewItem()
	expectedItem2 := factories.NewItem()
	expectedGiftProduct := factories.NewProduct()
	expectedGiftProduct.IsGift = true
	expectedItems := []domain.Item{*expectedItem1, *expectedItem2}
	expectedCart := &domain.Cart{
		Items: expectedItems,
	}
	expectedCart.CalculateTotals()

	s.productRepo.On(findAllMethod, cart.Items).Return(expectedItems, nil)
	s.productRepo.On(findLastByIsGiftMethod, true).Return(domain.Product{}, errors.New(""))
	s.discountUsecase.On(calculateDiscountsMethod, expectedItems).Return(expectedItems, nil)
	s.config.On(isBlackFridayActiveMethod).Return(true)
	s.config.On(getBlackFridayDayMethod).Return(time.Now().Day())
	s.config.On(getBlackFridayMonthMethod).Return(time.Now().Month())

	err := s.useCase.Checkout(cart)

	assert.Exactly(s.T(), expectedCart, cart)
	assert.Nil(s.T(), err)
	s.productRepo.AssertNumberOfCalls(s.T(), findAllMethod, 1)
	s.productRepo.AssertNumberOfCalls(s.T(), findLastByIsGiftMethod, 1)
	s.discountUsecase.AssertNumberOfCalls(s.T(), calculateDiscountsMethod, 1)
}

func (s *CheckoutUseCaseSuite) TestCheckoutItemsNotFound() {
	cart := factories.NewCart()
	expectedError := &customErr.NotFound{Entity: "Product", Err: errors.New("the following products were not found: " + string(cart.Items[0].Product.ID))}

	s.productRepo.On(findAllMethod, cart.Items).Return(nil, expectedError)

	err := s.useCase.Checkout(cart)

	assert.Exactly(s.T(), expectedError, err)
	s.productRepo.AssertNumberOfCalls(s.T(), findAllMethod, 1)
	s.discountUsecase.AssertNumberOfCalls(s.T(), calculateDiscountsMethod, 0)
}

func (s *CheckoutUseCaseSuite) TestCheckoutMoreThanAllowedGifts() {
	cart := factories.NewCart()
	expectedItem1 := factories.NewItem()
	expectedItem1.Product.IsGift = true
	expectedItem2 := factories.NewItem()
	expectedItem2.Product.IsGift = true
	expectedItems := []domain.Item{*expectedItem1, *expectedItem2}

	expectedError := &customErr.NotValid{Input: "Gift Items", Err: errors.New("more than allowed gifts")}

	s.productRepo.On(findAllMethod, cart.Items).Return(expectedItems, nil)
	s.config.On(getAllowedGiftsPerCartMethod).Return(1)

	err := s.useCase.Checkout(cart)

	assert.Exactly(s.T(), expectedError, err)
	s.productRepo.AssertNumberOfCalls(s.T(), findAllMethod, 1)
	s.discountUsecase.AssertNumberOfCalls(s.T(), calculateDiscountsMethod, 0)
}

func (s *CheckoutUseCaseSuite) TestCheckoutDiscountError() {
	cart := factories.NewCart()
	expectedItem1 := factories.NewItem()
	expectedItem2 := factories.NewItem()
	expectedItems := []domain.Item{*expectedItem1, *expectedItem2}
	expectedCart := &domain.Cart{
		Items: expectedItems,
	}
	expectedCart.CalculateTotals()

	s.productRepo.On(findAllMethod, cart.Items).Return(expectedItems, nil)
	s.discountUsecase.On(calculateDiscountsMethod, expectedItems).Return(expectedItems, errors.New(""))
	s.config.On(isBlackFridayActiveMethod).Return(false)

	err := s.useCase.Checkout(cart)

	assert.Nil(s.T(), err) // should not return the error
	assert.Exactly(s.T(), expectedCart, cart)
	s.productRepo.AssertNumberOfCalls(s.T(), findAllMethod, 1)
	s.discountUsecase.AssertNumberOfCalls(s.T(), calculateDiscountsMethod, 1)
}
