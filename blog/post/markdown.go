package post

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	rendererHTML "github.com/yuin/goldmark/renderer/html"
)

// Errors
var (
	ErrInvalidType      = errors.New("Invalid type while converting from interface")
	ErrPostNotPublished = errors.New("Post was not published yet")
)

type markdown struct {
	parser goldmark.Markdown
	isDev  bool
}

// New returns a new markdown service
func New(isDev bool) Service {
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

	return &markdown{md, isDev}
}

func (m *markdown) GetAllPosts() (*[]Post, error) {
	files, err := ioutil.ReadDir("./static/markdown")
	if err != nil {
		return &[]Post{}, err
	}

	var posts []Post
	for _, f := range files {
		p, err := m.buildPost(
			fmt.Sprintf(
				"./static/markdown/%s",
				f.Name(),
			),
		)
		if err == ErrPostNotPublished {
			continue
		}

		if err != nil {
			return &posts, err
		}

		posts = append(posts, *p)
	}

	return &posts, nil
}

func (m *markdown) buildPost(filePath string) (*Post, error) {
	context := parser.NewContext()

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &Post{}, err
	}

	var buf bytes.Buffer
	err = m.parser.Convert(b, &buf, parser.WithContext(context))
	if err != nil {
		return &Post{}, err
	}

	params := meta.Get(context)

	var published string
	if _, ok := params["Published"]; ok != false {
		published, ok = params["Published"].(string)
		if ok != true {
			return &Post{}, ErrInvalidType
		}
	}

	if published == "" && m.isDev == false {
		return &Post{}, ErrPostNotPublished
	}

	title, ok := params["Title"].(string)
	if ok != true {
		return &Post{}, ErrInvalidType
	}

	slug, ok := params["Slug"].(string)
	if ok != true {
		return &Post{}, ErrInvalidType
	}

	var image string
	if _, ok := params["Image"]; ok != false {
		image, ok = params["Image"].(string)
		if ok != true {
			return &Post{}, ErrInvalidType
		}
	}

	var summary string
	if _, ok := params["Summary"]; ok != false {
		summary, ok = params["Summary"].(string)
		if ok != true {
			return &Post{}, ErrInvalidType
		}
	}

	tags, err := castTagsToString(params)
	if err != nil {
		return &Post{}, err
	}

	p := &Post{
		Title:   title,
		Slug:    slug,
		Summary: summary,
		Image:   image,
		HTML:    template.HTML(buf.String()),
		Tags:    tags,
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
		return []string{}, ErrInvalidType
	}

	strTags := make([]string, 0, len(tags))

	for _, tag := range tags {
		t, ok := tag.(string)
		if ok == false {
			continue
		}

		strTags = append(strTags, t)
	}

	return strTags, nil
}
