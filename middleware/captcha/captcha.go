package captcha

import (
	"strings"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
)

type Config struct {
	SkipLogging bool
	CaptchaPath string
	StdWidth    int
	StdHeight   int
}

var (
	DefaultConfig = Config{
		SkipLogging: true,
		CaptchaPath: "/captcha/",
		StdWidth:    captcha.StdWidth,
		StdHeight:   captcha.StdHeight,
	}
)

func Captcha(conf Config) echo.MiddlewareFunc {
	//Defaults
	if conf.CaptchaPath == "" {
		conf.CaptchaPath = DefaultConfig.CaptchaPath
	}
	if conf.StdWidth == 0 {
		conf.StdWidth = DefaultConfig.StdWidth
	}
	if conf.StdHeight == 0 {
		conf.StdHeight = DefaultConfig.StdHeight
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			method := request.Method
			if method != "GET" {
				return next(c)
			}

			url := request.URL.Path
			if strings.HasPrefix(url, conf.CaptchaPath) {
				if !conf.SkipLogging {
					c.Logger().Debugf("Captcha server, url:%v", url)
				}
				//@TODO 验证码尺寸支持Query参数
				captcha.Server(conf.StdWidth, conf.StdHeight).ServeHTTP(c.Response(), c.Request())
				return nil
			} else {
				return next(c)
			}

		}
	}
}
