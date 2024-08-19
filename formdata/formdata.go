package formdata

import (
	"fmt"
	"strings"
)

const indent = "  "

type data struct {
	Body        string
	CodeSnippet string
	Language    string
	Tags        string
	Title       string
	Url         string
	UrlTitle    string
}

func New() *data {
	return &data{}
}

func (d *data) SetBody(body string) {
	d.Body = body
}

func (d *data) SetCodeSnippet(codeSnippet string) {
	d.CodeSnippet = codeSnippet
}
func (d *data) SetLanguage(language string, _ int) {
	d.Language = language
}

func (d *data) SetTags(tags string) {
	d.Tags = tags
}

func (d *data) SetTitle(title string) {
	d.Title = title
}

func (d *data) SetUrl(url string) {
	d.Url = url
}

func (d *data) SetUrlTitle(urlTitle string) {
	d.UrlTitle = urlTitle
}

func (d *data) String() string {
	res := fmt.Sprintf("- [[%s]]\n", d.Title)
	if d.Body != "" {
		res += fmt.Sprintf("%s%s\n", indent, strings.Replace(d.Body, "\n", "\n"+indent, -1))
	}
	if d.Tags != "" {
		res += fmt.Sprintf("%s%s\n", indent, d.Tags)
	}
	if d.Language != "" {
		res += fmt.Sprintf("%s[[%s]]\n", indent, d.Language)
	}
	if d.CodeSnippet != "" {
		res += fmt.Sprintf("%s```%s\n%s\n%s```\n", indent, d.Language, d.CodeSnippet, indent)
	}
	if d.Url != "" {
		res += fmt.Sprintf("%s[%s](%s)\n", indent, d.UrlTitle, d.Url)
	}

	return res
}
