package page

import (
	"bytes"
	"html/template"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type markdown struct {
	parser goldmark.Markdown
}

// Markdown interface
type Markdown interface {
	MarkdownToHTML(content []byte) (template.HTML, error)
	BuildPage(content []byte) (*Page, error)
}

// NewMarkdown returns a new markdown service
func NewMarkdown() Markdown {
	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	return &markdown{md}
}

func (m *markdown) MarkdownToHTML(content []byte) (template.HTML, error) {
	context := parser.NewContext()

	var buf bytes.Buffer
	err := m.parser.Convert(content, &buf, parser.WithContext(context))
	if err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}

func (m *markdown) BuildPage(content []byte) (*Page, error) {
	context := parser.NewContext()

	var buf bytes.Buffer
	err := m.parser.Convert(content, &buf, parser.WithContext(context))
	if err != nil {
		return &Page{}, err
	}

	params := meta.Get(context)
	title, ok := params["Title"].(string)
	if ok != true {
		return &Page{}, err
	}

	slug, ok := params["Slug"].(string)
	if ok != true {
		return &Page{}, err
	}

	layout, ok := params["Layout"].(string)
	if ok != true {
		return &Page{}, err
	}

	var image string
	if _, ok := params["Image"]; ok != false {
		image, ok = params["Image"].(string)
		if ok != true {
			return &Page{}, err
		}
	}

	var summary string
	if _, ok := params["Summary"]; ok != false {
		summary, ok = params["Summary"].(string)
		if ok != true {
			return &Page{}, err
		}
	}

	var published string
	if _, ok := params["Published"]; ok != false {
		published, ok = params["Published"].(string)
		if ok != true {
			return &Page{}, err
		}
	}

	p := &Page{
		MetaTitle: title,
		Title:     title,
		Slug:      slug,
		Summary:   summary,
		Image:     image,
		Layout:    layout,
		HTML:      template.HTML(buf.String()),
		Published: published,
	}

	return p, nil
}
