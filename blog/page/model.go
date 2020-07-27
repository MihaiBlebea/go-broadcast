package page

import (
	"html/template"
	"io"
	"time"
)

// Page model
type Page struct {
	HTML      template.HTML
	Params    interface{}
	MetaTitle string
	Image     string
	Title     string
	Slug      string
	Summary   string
	Tags      []string
	Template  *template.Template
	Layout    string
	Published string
	Articles  []*Page
}

// Render the page
func (p *Page) Render(w io.Writer) error {
	err := p.Template.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}

// GetPublished returns the publish date
func (p *Page) GetPublished() (time.Time, error) {
	timeFormat := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(timeFormat, p.Published)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
