package socket

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"

	"github.com/foxiswho/shop-go/module/log"
)

func socketHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			if len(msg) > 0 {
				log.Debugf("socket msg:" + msg)
			} else {
				break
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
