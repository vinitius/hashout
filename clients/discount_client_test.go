package clients

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"viniti.us/hashout/config/cb"
	"viniti.us/hashout/config/log"
	domain "viniti.us/hashout/models/checkout"
	discountPB "viniti.us/hashout/pb"
	"viniti.us/hashout/test/factories"
	mocks "viniti.us/hashout/test/mocks/pb"
)

const (
	getDiscountMethod = "GetDiscount"
)

type DiscountClientSuite struct {
	suite.Suite
	grpc *mocks.DiscountClient
	ctx  context.Context
	cli  DiscountClient
}

func (s *DiscountClientSuite) SetupTest() {
	log.SetupLogger()
	s.grpc = new(mocks.DiscountClient)
	s.ctx = context.Background()
	s.cli = DiscountClient{grpc: s.grpc, ctx: s.ctx, command: cb.NewDiscountCB()}
}

func TestRun(t *testing.T) {
	s := new(DiscountClientSuite)
	suite.Run(t, s)
}

func (s *DiscountClientSuite) TestGetDiscountSuccessfully() {
	item := factories.NewItem()
	expectedRate := 0.10
	expectedItem := &domain.Item{
		Product:      item.Product,
		Quantity:     item.Quantity,
		DiscountRate: float32(expectedRate) * 100,
		UnitAmount:   item.UnitAmount,
		TotalAmount:  item.TotalAmount,
	}

	s.grpc.On(getDiscountMethod, s.ctx, &discountPB.GetDiscountRequest{ProductID: item.Product.ID}).Return(&discountPB.GetDiscountResponse{Percentage: float32(expectedRate)}, nil)

	err := s.cli.GetDiscount(item)

	assert.Exactly(s.T(), expectedItem, item)
	assert.Nil(s.T(), err)
	s.grpc.AssertNumberOfCalls(s.T(), getDiscountMethod, 1)
}

func (s *DiscountClientSuite) TestGetDiscountError() {
	item := factories.NewItem()
	item.DiscountRate = 0

	expectedError := errors.New("")

	s.grpc.On(getDiscountMethod, s.ctx, &discountPB.GetDiscountRequest{ProductID: item.Product.ID}).Return(nil, expectedError)

	err := s.cli.GetDiscount(item)

	assert.Exactly(s.T(), expectedError, err)
	assert.Equal(s.T(), float32(0), item.DiscountRate)
	s.grpc.AssertNumberOfCalls(s.T(), getDiscountMethod, 1)
}
