package formdata

import (
	"fmt"
	"strings"
)

const indent = "  "

type Field struct {
	value   string
	include bool
}

type Data struct {
	Body        Field
	CodeSnippet string
	Language    string
	Tags        string
	Title       string
	Url         string
	UrlTitle    string
}

func New() *Data {
	return &Data{}
}

func (d *Data) SetBody(body string) {
	d.Body.value = body
}

func (d *Data) SetCodeSnippet(codeSnippet string) {
	d.CodeSnippet = codeSnippet
}
func (d *Data) SetLanguage(language string, _ int) {
	d.Language = language
}

func (d *Data) SetTags(tags string) {
	d.Tags = tags
}

func (d *Data) SetTitle(title string) {
	d.Title = title
}

func (d *Data) SetUrl(url string) {
	d.Url = url
}

func (d *Data) SetUrlTitle(urlTitle string) {
	d.UrlTitle = urlTitle
}

func (d *Data) PrintBody() string {
	if d.Body.include && d.Body.value != "" {
		return fmt.Sprintf("%s%s\n", indent, strings.Replace(d.Body.value, "\n", "\n"+indent, -1))
	}
	return ""
}

func (d *Data) String() string {
	res := fmt.Sprintf("- [[%s]]\n", d.Title)
	res += d.PrintBody()
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
