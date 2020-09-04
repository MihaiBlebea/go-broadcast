package page

import (
	"fmt"
	"html/template"
	"io"
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
