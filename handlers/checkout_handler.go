package handlers

import (
	"viniti.us/hashout/usecase/checkout"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	useCase checkout.Service
}

func NewCheckoutHandler(useCase checkout.Service) CheckoutHandler {
	return CheckoutHandler{useCase: useCase}
}

func (h CheckoutHandler) Routes(r *gin.Engine) {
	r.POST("/checkout", h.checkout)
}

func (h CheckoutHandler) checkout(c *gin.Context) {

}
