package main

import (
	"fmt"
	"github.com/foxiswho/shop-go/middleware/cache"
	"github.com/labstack/echo"
	"time"
)

func main() {
	r := echo.Default()

	store := cache.NewInMemoryStore(time.Second)
	// Cached Page
	r.GET("/ping", func(c *echo.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	r.GET("/cache_ping", cache.CachePage(store, time.Minute, func(c *echo.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	}))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
