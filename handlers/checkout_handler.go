package handlers

import (
	"net/http"

	dto "viniti.us/hashout/handlers/dto/checkout"
	customErr "viniti.us/hashout/models/errors"
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
	var req dto.Checkout
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(&customErr.NotValid{Input: "Checkout", Err: err})
		return
	}

	cart, err := req.ToDomain()
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.useCase.Checkout(&cart); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.ToCheckoutResponse(cart))
}
