package handlers

import (
	"order-management/config/log"

	"viniti.us/hashout/models/customErr"

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

func fromMalFormed(err *customErr.MalFormed) *ApiError {
	return &ApiError{
		Code:    400,
		Message: err.Error(),
		Reason:  "Malformed Payload.",
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
			case *expectedErrors.NotFound:
				apiError = fromNotFound(parsed)
			case *expectedErrors.NotValid:
				apiError = fromNotValid(parsed)
			case *expectedErrors.MalFormed:
				apiError = fromMalFormed(parsed)
			default:
				apiError = fromGeneric(parsed)
			}

			c.AbortWithStatusJSON(apiError.Code, apiError)
		}
	}
}
