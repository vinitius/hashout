package app

import (
	"time"

	"github.com/spf13/viper"
)

type Config interface {
	GetAllowedGiftsPerCart() int
	IsBlackFridayGiftActive() bool
	GetBlackFridayDay() int
	GetBlackFridayMonth() time.Month
}

type HashoutApp struct {
	Config
}

func (a HashoutApp) GetAllowedGiftsPerCart() int {
	return viper.GetInt("ALLOWED_GIFTS_PER_CART")
}

func (a HashoutApp) IsBlackFridayGiftActive() bool {
	return viper.GetBool("BLACK_FRIDAY_GIFT_TOGGLE")
}

func (a HashoutApp) GetBlackFridayDay() int {
	return 26
}

func (a HashoutApp) GetBlackFridayMonth() time.Month {
	return time.November
}

func NewAppConfig() HashoutApp {
	return HashoutApp{}
}
