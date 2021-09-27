package checkout

import (
	"time"

	"viniti.us/hashout/config/log"
	"viniti.us/hashout/models/checkout"
	"viniti.us/hashout/usecase/discounts"
)

const (
	BlackFridayDay   int        = 26
	BlackFridayMonth time.Month = time.November
)

type UseCase struct {
	repo             Repository
	discountsUseCase discounts.Service
}

func NewUseCase(repo Repository, discountsUseCase discounts.Service) UseCase {
	return UseCase{repo: repo, discountsUseCase: discountsUseCase}
}

func (u UseCase) Checkout(c *checkout.Cart) error {
	products, err := u.repo.FindAll(c.Items)
	if err != nil {
		return err
	}

	c.Items = products

	if u.IsBlackFriday() {
		gift, err := u.repo.FindLastByIsGift(true)
		if err != nil {
			return err
		}
		c.AddGift(gift)
	}

	withDiscounts, err := u.discountsUseCase.CalculateDiscounts(products)
	if err != nil {
		log.Logger.Warn("error upon calculating discounts: ", err.Error())
		log.Logger.Info("moving on with no discounts!")
	} else {
		c.Items = withDiscounts
	}

	c.CalculateTotals()

	return nil
}

func (u UseCase) IsBlackFriday() bool {
	_, currentMonth, currentDay := time.Now().Date()
	return BlackFridayMonth == currentMonth && BlackFridayDay == currentDay
}
