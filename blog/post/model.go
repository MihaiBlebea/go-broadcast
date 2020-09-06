package post

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

const format = "2006-01-02 15:04:05"

// Post model
type Post struct {
	HTML      template.HTML
	Image     string
	Title     string
	Slug      string
	Summary   string
	Tags      []string
	Published time.Time
}

// SetPublished set publish time from string
func (p *Post) SetPublished(date string) error {
	published, err := time.Parse(format, date)
	if err != nil {
		return err
	}

	p.Published = published

	return nil
}

// GetFormatPublished returns formatted published date
func (p *Post) GetFormatPublished() string {
	return fmt.Sprintf(
		"%d %s %d",
		p.Published.Day(),
		p.Published.Month().String()[0:3],
		p.Published.Year(),
	)
}

// GetShareOnTwitterLink returns a link pre-polulated from the content of the page
func (p *Post) GetShareOnTwitterLink() string {
	slug := fmt.Sprintf(
		"article/%s",
		p.Slug,
	)

	tags := []string{"MBlebea"}
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
func (p *Post) GetShareOnFacebookLink() string {
	slug := fmt.Sprintf(
		"article/%s",
		p.Slug,
	)

	return fmt.Sprintf(
		"https://www.facebook.com/sharer/sharer.php?u=https://mihaiblebea.com/%s",
		slug,
	)
}

// GetShareOnLinkedinLink returns a link for sharing the post on linkedin
func (p *Post) GetShareOnLinkedinLink() string {
	slug := fmt.Sprintf(
		"article/%s",
		p.Slug,
	)

	return fmt.Sprintf(
		"https://www.linkedin.com/shareArticle?mini=true&url=https://mihaiblebea.com/%s&title=%s&summary=%s&source=",
		slug,
		p.Title,
		p.Summary,
	)
}

// IsDraft returns false if the publish date if before today
func (p *Post) IsDraft() bool {
	if p.Published.IsZero() {
		return true
	}

	return false
}
