package render

import (
	"bytes"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"

	. "github.com/foxiswho/echo-go/conf"
	"github.com/foxiswho/echo-go/module/log"
	bdTmpl "github.com/foxiswho/echo-go/template"
)

func pongo2() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}

			tmpl, context, err := getContext(c)
			if err == nil {
				c.Render(http.StatusOK, tmpl+Conf.Tmpl.Suffix, context)
			} else {
				log.Debugf("Pongo2 render Error: %v", err)
			}

			return nil
		}
	}
}

type BindataFileLoader struct {
	baseDir string
}

func (bf BindataFileLoader) Abs(base, name string) string {
	_, exist := bdTmpl.AssetInfo(name)
	if exist == nil {
		return name
	}

	// Our own base dir has always priority; if there's none
	// we use the path provided in base.
	if base != "" {
		return filepath.Join(filepath.Dir(base), name)
	}

	return filepath.Join(bf.baseDir, name)
}

func (bf BindataFileLoader) Get(path string) (io.Reader, error) {

	buf, err := bdTmpl.Asset(path)
	if err != nil {
		log.Debugf("Pongo2 bindata file load err: %v", err)
		return nil, err
	}

	return bytes.NewReader(buf), nil
}
