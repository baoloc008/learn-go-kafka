package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"learn-go-kafka/kafka"
	"net/http"
)

func PushMessage(c echo.Context) error {
	ctx := context.Background()
	defer ctx.Done()
	form := &struct {
		Text string `form:"text" json:"text"`
	}{}
	err := c.Bind(form)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("error while binding data: %s", err.Error()))
	}
	formInBytes, err := json.Marshal(form)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("error while marshalling json: %s", err.Error()))
	}
	err = kafka.Write(ctx, nil, formInBytes)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, fmt.Sprintf("error while push message into kafka: %s", err.Error()))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    1,
		"message": "success push data into kafka",
		"data":    form,
	})
}
