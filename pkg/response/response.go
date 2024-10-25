package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type responseOpts func(*response) error

type response struct {
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WithMessage(message string) responseOpts {
	return func(r *response) error {
		r.Message = message

		return nil
	}
}

func WithData(data any) responseOpts {
	return func(r *response) error {
		r.Data = data

		return nil
	}
}

func GenerateResponse(c echo.Context, code int, options ...responseOpts) error {
	// Default response without any configuration
	response := response{
		Code:    code,
		Message: http.StatusText(code),
		Data:    nil,
	}

	// Check for any custom configuration
	// Return error if there is any mistaken on configure the response
	for _, opt := range options {
		// Modify the original default response struct
		err := opt(&response)
		if err != nil {
			panic(err)
		}
	}

	return c.JSON(code, response)
}
