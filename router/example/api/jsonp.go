package api


import (
	"github.com/foxiswho/shop-go/router/base"
	"time"
	"math/rand"
	"net/http"
)


func JsonpHandler(c *base.BaseContext) error {
	callback := c.QueryParam("callback")
	var content struct {
		Response  string    `json:"response"`
		Timestamp time.Time `json:"timestamp"`
		Random    int       `json:"random"`
	}
	content.Response = "Sent via JSONP"
	content.Timestamp = time.Now().UTC()
	content.Random = rand.Intn(1000)

	return c.JSONP(http.StatusOK, callback, &content)
}