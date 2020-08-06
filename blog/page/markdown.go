package page

import (
	"bytes"
	"errors"
	"html/template"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	rendererHTML "github.com/yuin/goldmark/renderer/html"
)

// Errors
var (
	ErrInvalidType = errors.New("Invalid type while converting from interface")
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
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					html.WithLineNumbers(true),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			rendererHTML.WithHardWraps(),
			rendererHTML.WithXHTML(),
			rendererHTML.WithUnsafe(),
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
		return &Page{}, ErrInvalidType
	}

	slug, ok := params["Slug"].(string)
	if ok != true {
		return &Page{}, ErrInvalidType
	}

	layout, ok := params["Layout"].(string)
	if ok != true {
		return &Page{}, ErrInvalidType
	}

	var image string
	if _, ok := params["Image"]; ok != false {
		image, ok = params["Image"].(string)
		if ok != true {
			return &Page{}, ErrInvalidType
		}
	}

	var summary string
	if _, ok := params["Summary"]; ok != false {
		summary, ok = params["Summary"].(string)
		if ok != true {
			return &Page{}, ErrInvalidType
		}
	}

	var published string
	if _, ok := params["Published"]; ok != false {
		published, ok = params["Published"].(string)
		if ok != true {
			return &Page{}, ErrInvalidType
		}
	}

	var kind kind
	k, ok := params["Kind"].(string)
	if ok != true {
		return &Page{}, ErrInvalidType
	}

	switch k {
	case "article":
		kind = kindArticle
	default:
		kind = kindPage
	}

	tags, err := castTagsToString(params)
	if err != nil {
		return &Page{}, err
	}

	p := &Page{
		MetaTitle: title,
		Title:     title,
		Slug:      slug,
		Summary:   summary,
		Image:     image,
		Layout:    layout,
		HTML:      template.HTML(buf.String()),
		Tags:      tags,
		Kind:      kind,
	}

	p.SetPublished(published)

	return p, nil
}

func castTagsToString(params map[string]interface{}) ([]string, error) {
	if _, ok := params["Tags"]; ok == false {
		return []string{}, nil
	}

	tags, ok := params["Tags"].([]interface{})
	if ok == false {
		return []string{}, errors.New("Could not convert interface to slice of interfaces")
	}

	strTags := make([]string, len(tags))

	for _, tag := range tags {
		t, ok := tag.(string)
		if ok == false {
			return []string{}, errors.New("Could not convert interface to string")
		}

		strTags = append(strTags, t)
	}

	return strTags, nil
}
