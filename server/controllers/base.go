package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/jirayutcc/broadcast"
)

type Base interface {
	CreateTransaction(c echo.Context) error
	Monitor(ctx echo.Context) error
}

type BaseImpl struct {
	apiConnector *broadcast.Base
}

func InitController(ac *broadcast.Base) Base {
	return BaseImpl{
		apiConnector: ac,
	}
}

func (c BaseImpl) CreateTransaction(ctx echo.Context) error {
	var message broadcast.BroadcastTransaction

	err := ctx.Bind(&message)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	response, err := (*c.apiConnector).Broadcast(&message)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c BaseImpl) Monitor(ctx echo.Context) error {
	var message broadcast.GetTransaction

	hashMessage := ctx.Param("hash")
	message.TxHash = hashMessage

	response, err := (*c.apiConnector).Monitor(&message)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}
