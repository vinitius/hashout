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

// Checkout godoc
// @Summary checkout items within a cart.
// @Tags Cart
// @Produce  json
// @Param filters body dto.Checkout true "Products"
// @Success 200 {object} dto.CheckoutResponse
// @Failure 400 {object} ApiError "Invalid Or Missing Products"
// @Failure 404 {object} ApiError "Products Not Found"
// @Failure 500 {object} ApiError "Unexpected"
// @Failure default {object} ApiError
// @Router /checkout [post]
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
