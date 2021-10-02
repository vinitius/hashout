package checkout

import (
	"errors"
	"time"

	"github.com/spf13/viper"
	"viniti.us/hashout/config/log"
	"viniti.us/hashout/models/checkout"
	customErr "viniti.us/hashout/models/errors"
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

func (u UseCase) Checkout(c *checkout.Cart) (err error) {
	c.Items, err = u.repo.FindAll(c.Items)
	if err != nil {
		return err
	}

	contains, count := c.ContainsGift()
	if contains && count > viper.GetInt32("ALLOWED_GIFTS_PER_CART") {
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
	return viper.GetBool("BLACK_FRIDAY_GIFT_TOGGLE") && BlackFridayMonth == currentMonth && BlackFridayDay == currentDay
}
