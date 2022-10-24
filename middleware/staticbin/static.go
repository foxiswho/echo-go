// Based on https://github.com/olebedev/staticbin
package staticbin

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Options struct {
	// SkipLogging will disable [Static] log messages when a static file is served.
	SkipLogging bool
	// IndexFile defines which file to serve as index if it exists.
	IndexFile string
	// Path prefix
	Dir string
}

func (o *Options) init() {
	if o.IndexFile == "" {
		o.IndexFile = "index.html"
	}
}

// Static returns a middleware handler that serves static files in the given directory.
func Static(asset func(string) ([]byte, error), options ...Options) echo.MiddlewareFunc {
	if asset == nil {
		panic("asset is nil")
	}

	opt := Options{}
	for _, o := range options {
		opt = o
		break
	}
	opt.init()

	modtime := time.Now()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			method := request.Method
			if method != "GET" && method != "HEAD" {
				// Request is not correct. Go farther.
				// return echo.NewHTTPError(http.StatusMethodNotAllowed)
				return next(c)
			}

			url := request.URL.Path
			if !strings.HasPrefix(url, opt.Dir) {
				// return echo.NewHTTPError(http.StatusUnsupportedMediaType)
				return next(c)
			}
			file := strings.TrimPrefix(
				strings.TrimPrefix(url, opt.Dir),
				"/",
			)
			b, err := asset(file)

			if err != nil {
				// Try to serve the index file.
				b, err = asset(path.Join(file, opt.IndexFile))

				if err != nil {
					// Go farther if the asset could not be found.
					return next(c)
				}
			}

			if !opt.SkipLogging {
				c.Logger().Debugf("Static file, url:%v", url)
			}

			//缺少contentType MIME类型有BUG
			// return c.Blob(200, "", b)

			http.ServeContent(c.Response(), c.Request(), url, modtime, bytes.NewReader(b))
			return nil
		}
	}
}
