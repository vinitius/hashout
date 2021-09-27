package handlers

import (
	"viniti.us/hashout/config/log"

	customErr "viniti.us/hashout/models/errors"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
} // @name checkoutApiError

func fromNotFound(err *customErr.NotFound) *ApiError {
	return &ApiError{
		Code:    404,
		Message: err.Error(),
		Reason:  "Requested Entity was Not Found.",
	}

}

func fromNotValid(err *customErr.NotValid) *ApiError {
	return &ApiError{
		Code:    400,
		Message: err.Error(),
		Reason:  "Invalid Payload or Parameters.",
	}
}

func fromGeneric(err error) *ApiError {
	return &ApiError{
		Code:    500,
		Message: err.Error(),
		Reason:  "Unexpected Server Error. Please, check logs.",
	}
}

func ApiErrors() gin.HandlerFunc {
	return apiErrorReporter(gin.ErrorTypeAny)
}

func apiErrorReporter(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)
		var apiError *ApiError

		if len(detectedErrors) > 0 {
			log.Logger.Warn("ApiError: Detected: " + detectedErrors.Errors()[0])
			err := detectedErrors[0].Err
			switch parsed := err.(type) {
			case *customErr.NotFound:
				apiError = fromNotFound(parsed)
			case *customErr.NotValid:
				apiError = fromNotValid(parsed)
			default:
				apiError = fromGeneric(parsed)
			}

			c.AbortWithStatusJSON(apiError.Code, apiError)
		}
	}
}
