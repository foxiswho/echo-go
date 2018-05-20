package goods

import (
	"github.com/foxiswho/echo-go/router/base"
	"github.com/foxiswho/echo-go/service/goods_service"
	"net/http"
)

func CreateHandler(c *base.BaseContext) error {
	goods_service.CreateGoodsxxxxxx()
	c.HTML(http.StatusOK, "SUCCESS")
	return nil
}
