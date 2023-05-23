package helpers

import "github.com/labstack/echo/v4"

type ResponseModel struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func Response(e echo.Context, status_code int, response_data ResponseModel) error {
	return e.JSON(status_code, response_data)
}
