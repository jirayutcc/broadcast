package core

import (
	"github.com/labstack/echo/v4"
	"github.com/jirayutcc/broadcast"
	"github.com/jirayutcc/transaction-broadcast/controllers"
)

func Setup() *echo.Echo {
	e := echo.New()

	ac := broadcast.InitBaseAPI()
	c := controllers.InitController(&ac)

	e.GET("broadcast/:hash", c.Monitor)
	e.POST("broadcast", c.CreateTransaction)

	return e
}
