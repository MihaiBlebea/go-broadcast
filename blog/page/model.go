package page

import (
	"fmt"
	"html/template"
	"io"
	"reflect"
)

// Page model
type Page struct {
	Params       interface{}
	Template     *template.Template
	TemplateName string
}

// Render the page
func (p *Page) Render(w io.Writer) error {
	file := fmt.Sprintf("%s.gohtml", p.TemplateName)

	return p.Template.ExecuteTemplate(w, file, p)
}

// IsSet check if attribute is set on Params interface
func (p *Page) IsSet(name string) bool {
	v := reflect.ValueOf(p.Params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	return v.FieldByName(name).IsValid()
}
