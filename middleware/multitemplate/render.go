// Base on github.com/gin-gonic/contrib/renders/multitemplate
package multitemplate

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

/**
 *
 */
type Render map[string]*template.Template

func New() Render {
	return make(Render)
}

func (r Render) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	r[name] = tmpl
}

func (r Render) AddFromFiles(name string, files ...string) *template.Template {
	tmpl := template.Must(template.ParseFiles(files...))
	r.Add(name, tmpl)
	return tmpl
}

func (r Render) AddFromGlob(name, glob string) *template.Template {
	tmpl := template.Must(template.ParseGlob(glob))
	r.Add(name, tmpl)
	return tmpl
}

func (r *Render) AddFromString(name, templateString string) *template.Template {
	tmpl := template.Must(template.New("").Parse(templateString))
	r.Add(name, tmpl)
	return tmpl
}

func (r Render) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var t *template.Template
	t = r[name]
	return t.Execute(w, data)
}
