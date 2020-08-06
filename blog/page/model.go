package page

import (
	"fmt"
	"html/template"
	"io"
	"math"
	"strings"
	"time"
)

const format = "2006-01-02 15:04:05"

type kind int

const (
	kindArticle kind = iota
	kindPage
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
	Published time.Time
	Articles  []*Page
	Kind      kind
}

// Render the page
func (p *Page) Render(w io.Writer) error {
	err := p.Template.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}

// SetPublished set publish time from string
func (p *Page) SetPublished(date string) error {
	published, err := time.Parse(format, date)
	if err != nil {
		return err
	}

	p.Published = published

	return nil
}

// GetHumanReadablePublished returns a human readale representation of the published time
func (p *Page) GetHumanReadablePublished() string {
	duration := time.Now().Sub(p.Published)

	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"day", days},
		{"hour", hours},
	}

	parts := []string{}

	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}

	return fmt.Sprintf(
		"%s ago",
		strings.Join(parts, " "),
	)
}

// GetShareOnTwitterLink returns a link pre-polulated from the content of the page
func (p *Page) GetShareOnTwitterLink() string {
	slug := p.Slug
	if p.Kind == kindArticle {
		slug = fmt.Sprintf(
			"article/%s",
			p.Slug,
		)
	}

	tags := []string{"mihaiblebea"}
	for _, tag := range p.Tags {
		tags = append(tags, tag)
	}

	return fmt.Sprintf(
		"http://twitter.com/share?text=%s&url=https://mihaiblebea.com/%s&hashtags=%s",
		p.Summary,
		slug,
		strings.Join(tags, ","),
	)
}

// GetShareOnFacebookLink returns a link for sharring the post on facebook
func (p *Page) GetShareOnFacebookLink() string {
	slug := p.Slug
	if p.Kind == kindArticle {
		slug = fmt.Sprintf(
			"article/%s",
			p.Slug,
		)
	}

	return fmt.Sprintf(
		"https://www.facebook.com/sharer/sharer.php?u=https://mihaiblebea.com/%s",
		slug,
	)
}

// GetShareOnLinkedinLink returns a link for sharing the post on linkedin
func (p *Page) GetShareOnLinkedinLink() string {
	slug := p.Slug
	if p.Kind == kindArticle {
		slug = fmt.Sprintf(
			"article/%s",
			p.Slug,
		)
	}

	return fmt.Sprintf(
		"https://www.linkedin.com/shareArticle?mini=true&url=https://mihaiblebea.com/%s&title=%s&summary=%s&source=",
		slug,
		p.Title,
		p.Summary,
	)
}
