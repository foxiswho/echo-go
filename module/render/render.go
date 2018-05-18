package render

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"

	"github.com/foxiswho/shop-go/middleware/multitemplate"
	"github.com/foxiswho/shop-go/middleware/pongo2echo"
	"github.com/foxiswho/shop-go/middleware/session"

	. "github.com/foxiswho/shop-go/conf"
	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/log"
	MT "github.com/foxiswho/shop-go/template"
)

func Render() echo.MiddlewareFunc {
	if Conf.Tmpl.Type == PONGO2 {
		return pongo2()
	} else {
		return render()
	}
}

func render() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}

			tmpl, context, err := getContext(c)
			if err == nil {
				c.Render(http.StatusOK, tmpl, context)
			} else {
				c.Logger().Errorf("Render Error: %v, tmpl %v, content %v", err, tmpl, context)
			}

			return nil
		}
	}
}

func getContext(c echo.Context) (tmpl string, context map[string]interface{}, err error) {
	tmplName := c.Get("tmpl")
	tmplNameValue, isString := tmplName.(string)
	tmplData := c.Get("data")

	// 模板未定义
	if !isString {
		return "", nil, errors.New("No tmpl defined!")
	}

	// 公共模板数据
	commonDatas := getCommonContext(c)

	// 模板数据
	if tmplData != nil {
		contextData, isMap := tmplData.(map[string]interface{})

		if isMap {
			for key, value := range commonDatas {
				contextData[key] = value
			}

			return tmplNameValue, contextData, nil
		}
	}

	return tmplNameValue, commonDatas, nil

}

func getCommonContext(c echo.Context) map[string]interface{} {
	a := auth.Default(c)
	fmt.Println("a.Admin======>")
	fmt.Println("a.Admin======>")
	fmt.Println("a.Admin======>")
	fmt.Println("a.Admin======>")
	fmt.Println("a.Admin======>")
	fmt.Println("a.Admin======>")
	fmt.Println("a.Admin======>",a.User)
	// 公共模板数据
	commonDatas := make(map[string]interface{})
	//commonDatas["_user"] = a.Admin.(*model.Admin)
	commonDatas["_user"] = a.User

	// 配置
	commonDatas["_conf"] = Conf

	// @TODO i18n

	// CSRF
	csrf := c.Get("_csrf")
	commonDatas["_csrf"] = csrf

	// Error
	s := session.Default(c)
	commonDatas["_error"] = s.Flashes("_error")

	path := c.Request().URL.Path
	uri := c.Request().RequestURI

	// 登录、注册、退出页面取已有RedirectParam
	redirect := c.QueryParam(auth.RedirectParam)
	switch path {
	case "/login":
		uri = redirect
	case "/register":
		uri = redirect
	case "/logout":
		uri = redirect
	default:
	}

	c.Logger().Debugf("Path : %v", path)
	c.Logger().Debugf("URI : %v", uri)

	commonDatas["requestUrl"] = uri

	return commonDatas
}

/**
 * 模板加载
 * 支持文件/Bindata加载模板
 */

func LoadTemplates() echo.Renderer {
	switch Conf.Tmpl.Type {
	case PONGO2:
		switch Conf.Tmpl.Data {
		case BINDATA:
			return pongo2echo.New(
				pongo2echo.RenderOptions{
					TmplLoader: BindataFileLoader{baseDir: Conf.Tmpl.Dir},
					ContentType: "text/html; charset=utf-8",
					Debug:       !Conf.ReleaseMode,
				})
		default:
			return pongo2echo.New(
				pongo2echo.RenderOptions{
					TemplateDir: Conf.Tmpl.Dir,
					ContentType: "text/html; charset=utf-8",
					Debug:       !Conf.ReleaseMode,
				})
		}
	case TEMPLATE:
	default:
		switch Conf.Tmpl.Data {
		case BINDATA:
			return loadTemplatesBindata(Conf.Tmpl.Dir)
		default:
			return loadTemplatesDefault(Conf.Tmpl.Dir)
		}
	}

	return loadTemplatesDefault(Conf.Tmpl.Dir)
}

func loadTemplatesDefault(templateDir string) *multitemplate.Render {
	r := multitemplate.New()

	layoutDir := templateDir + "/layout/"
	layouts, err := filepath.Glob(layoutDir + "*/*" + Conf.Tmpl.Suffix)
	if err != nil {
		panic(err.Error())
	}

	includeDir := templateDir + "/include/"
	includes, err := filepath.Glob(includeDir + "*" + Conf.Tmpl.Suffix)
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		tmpl := template.Must(template.ParseFiles(files...))
		tmplName := strings.TrimPrefix(layout, layoutDir)
		tmplName = strings.TrimSuffix(tmplName, Conf.Tmpl.Suffix)
		log.Debugf("Tmpl add " + tmplName)
		r.Add(tmplName, tmpl)
	}
	return &r
}

func loadTemplatesBindata(templateDir string) *multitemplate.Render {
	r := multitemplate.New()

	layoutDir := templateDir + "/layout"
	layoutDirs, err := MT.AssetDir(layoutDir)
	if err != nil {
		panic(err.Error())
	}

	var layouts []string
	for _, dir := range layoutDirs {
		files, err := MT.AssetDir(layoutDir + "/" + dir)
		if err != nil {
			panic(err.Error())
		}

		// 过滤非.tmpl后缀模板
		layoutFiels, err := tmplsFilter(files, layoutDir+"/"+dir)
		if err != nil {
			panic(err.Error())
		}

		layouts = append(layouts, layoutFiels...)
	}

	includeDir := templateDir + "/include"
	includeFiels, err := MT.AssetDir(includeDir)
	if err != nil {
		panic(err.Error())
	}
	// 过滤非.tmpl后缀模板
	includes, err := tmplsFilter(includeFiels, includeDir)
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		tmpl := template.Must(parseBindataFiles(files...))
		tmplName := strings.TrimPrefix(layout, layoutDir+"/")
		tmplName = strings.TrimSuffix(tmplName, Conf.Tmpl.Suffix)
		log.Debugf("Tmpl add " + tmplName)
		r.Add(tmplName, tmpl)
	}
	return &r
}

// 过滤非tmpl后缀模板文件
func tmplsFilter(files []string, dir string) ([]string, error) {
	var tmpls []string
	for _, file := range files {
		if strings.HasSuffix(file, Conf.Tmpl.Suffix) {
			tmpls = append(tmpls, dir+"/"+file)
		}
	}
	return tmpls, nil
}

// parseFiles is the helper for the method and function. If the argument
// template is nil, it is created from the first file.
func parseBindataFiles(filenames ...string) (*template.Template, error) {
	var t *template.Template
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
	}
	for _, filename := range filenames {
		b, err := MT.Asset(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		name := filepath.Base(filename)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
