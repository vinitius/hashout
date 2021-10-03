package checkout

import (
	"errors"
	"time"

	"viniti.us/hashout/config/app"
	"viniti.us/hashout/config/log"
	"viniti.us/hashout/models/checkout"
	customErr "viniti.us/hashout/models/errors"
	"viniti.us/hashout/usecase/discounts"
)

type UseCase struct {
	repo             Repository
	discountsUseCase discounts.Service
	config           app.Config
}

func NewUseCase(repo Repository, discountsUseCase discounts.Service, config app.Config) UseCase {
	return UseCase{repo: repo, discountsUseCase: discountsUseCase, config: config}
}

func (u UseCase) Checkout(c *checkout.Cart) (err error) {
	c.Items, err = u.repo.FindAll(c.Items)
	if err != nil {
		return err
	}

	contains, count := c.ContainsGift()
	if contains && count > u.config.GetAllowedGiftsPerCart() {
		return &customErr.NotValid{Input: "Gift Items", Err: errors.New("more than allowed gifts")}
	}

	c.Items, err = u.discountsUseCase.CalculateDiscounts(c.Items)
	if err != nil {
		log.Logger.Warn("error upon calculating discounts for one or more products: ", err.Error())
	}

	if u.IsBlackFridayGiftActive() {
		gift, err := u.repo.FindLastByIsGift(true)
		if err != nil {
			log.Logger.Warn("error adding a gift product: ", err.Error())
		} else {
			c.AddGift(gift)
		}
	}

	c.CalculateTotals()

	return nil
}

func (u UseCase) IsBlackFridayGiftActive() bool {
	_, currentMonth, currentDay := time.Now().Date()
	return u.config.IsBlackFridayGiftActive() && u.config.GetBlackFridayMonth() == currentMonth && u.config.GetBlackFridayDay() == currentDay
}
