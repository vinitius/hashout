package storage

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
	mocks "viniti.us/hashout/test/mocks/db"
)

const (
	findProductByIDMethod         = "FindProductByID"
	findLastProductByIsGiftMethod = "FindLastProductByIsGift"
)

type ProductRepositorySuite struct {
	suite.Suite
	db   *mocks.ProductDataset
	repo ProductRepository
}

func (s *ProductRepositorySuite) SetupTest() {
	log.SetupLogger()
	s.db = new(mocks.ProductDataset)
	s.repo = ProductRepository{db: s.db}
}

func TestRun(t *testing.T) {
	s := new(ProductRepositorySuite)
	suite.Run(t, s)
}

func (s *ProductRepositorySuite) TestFindAllSuccessfully() {
	item1 := factories.NewItem()
	item1.Product.ID = 1
	item2 := factories.NewItem()
	item2.Product.ID = 2
	items := []domain.Item{*item1, *item2}

	foundProduct1 := factories.NewProduct()
	foundProduct2 := factories.NewProduct()

	expectedItem1 := domain.Item{
		Product:      *foundProduct1,
		Quantity:     item1.Quantity,
		DiscountRate: item1.DiscountRate,
		UnitAmount:   foundProduct1.Amount,
		TotalAmount:  foundProduct1.Amount * item1.Quantity,
		Discount:     0,
	}
	expectedItem2 := domain.Item{
		Product:      *foundProduct2,
		Quantity:     item2.Quantity,
		DiscountRate: item2.DiscountRate,
		UnitAmount:   foundProduct2.Amount,
		TotalAmount:  foundProduct2.Amount * item2.Quantity,
		Discount:     0,
	}

	expectedItems := []domain.Item{expectedItem1, expectedItem2}

	s.db.On(findProductByIDMethod, item1.Product.ID).Return(*foundProduct1, true)
	s.db.On(findProductByIDMethod, item2.Product.ID).Return(*foundProduct2, true)

	foundItems, err := s.repo.FindAll(items)

	assert.Exactly(s.T(), expectedItems, foundItems)
	assert.Nil(s.T(), err)
	s.db.AssertNumberOfCalls(s.T(), findProductByIDMethod, len(items))
}

func (s *ProductRepositorySuite) TestFindAllError() {
	item1 := factories.NewItem()
	items := []domain.Item{*item1}

	expectedError := &customErr.NotFound{Entity: "Product", Err: errors.New("the following products were not found: [" + strconv.Itoa(int(item1.Product.ID)) + "]")}

	s.db.On(findProductByIDMethod, item1.Product.ID).Return(domain.Product{}, false)

	foundItems, err := s.repo.FindAll(items)

	assert.Exactly(s.T(), expectedError, err)
	assert.Nil(s.T(), foundItems)
	s.db.AssertNumberOfCalls(s.T(), findProductByIDMethod, len(items))
}

func (s *ProductRepositorySuite) TestFindLastByIsGiftSuccessfully() {
	foundProduct1 := factories.NewProduct()

	s.db.On(findLastProductByIsGiftMethod, true).Return(*foundProduct1, true)

	found, err := s.repo.FindLastByIsGift(true)

	assert.Exactly(s.T(), foundProduct1, &found)
	assert.Nil(s.T(), err)
	s.db.AssertNumberOfCalls(s.T(), findLastProductByIsGiftMethod, 1)
}
