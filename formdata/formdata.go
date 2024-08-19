package formdata

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BottomleyIan/notes-tui/settings"
)

const indent = "  "

type Data struct {
	Body        string
	CodeSnippet string
	Date        string
	FieldData   map[string]string
	Language    string
	Title       string
	Url         string
	UrlTitle    string
	note        settings.Note
}

func New(note settings.Note) *Data {
	data := Data{}

	data.Date = time.Now().Format("2006-01-02")
	data.FieldData = make(map[string]string)
	for _, f := range note.Fields {
		data.FieldData[f.Name] = f.DefaultValue
	}
	data.note = note
	return &data
}

func (d *Data) SetBody(body string) {
	d.Body = body
}

func (d *Data) SetCodeSnippet(codeSnippet string) {
	d.CodeSnippet = codeSnippet
}

func (d *Data) SetDate(date string) {
	d.Date = date
}

func (d *Data) SetFieldData(fieldName, value string) {
	d.FieldData[fieldName] = value
}

func (d *Data) SetLanguage(language string, _ int) {
	d.Language = language
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

func (d *Data) PrintBody(sb *strings.Builder) {
	if d.Body != "" {
		sb.WriteString(fmt.Sprintf("%s%s\n", indent, strings.Replace(d.Body, "\n", "\n"+indent, -1)))
	}
}

func (d *Data) PrintFieldData(sb *strings.Builder) {
	for k, v := range d.FieldData {
		log.Println("PrintFieldData", k, v)
		sb.WriteString(fmt.Sprintf("%s%s:: %s\n", indent, k, FormatTags(v)))
	}
}

func FormatTags(tags string) string {
	if tags == "" {
		return ""
	}
	return "[[" + strings.Replace(tags, ",", "]][[", -1) + "]]"
}

func (d *Data) String() string {
	sb := strings.Builder{}

	if d.note.TagFirstLine {
		if d.Title == "" {

			sb.WriteString(fmt.Sprintf("\n- [[%s]] \n", d.note.Title))
		} else {
			sb.WriteString(fmt.Sprintf("\n- [[%s]] [[%s]]\n", d.note.Title, d.Title))
		}
	} else {

		sb.WriteString(fmt.Sprintf("\n- %s %s\n", d.note.Title, d.Title))
	}
	if d.Url != "" {
		sb.WriteString(fmt.Sprintf("%s link:: [%s](%s)\n", indent, d.UrlTitle, d.Url))
	}
	if d.Language != "" {
		sb.WriteString(fmt.Sprintf("%slanguage:: [[%s]]\n", indent, d.Language))
	}
	d.PrintFieldData(&sb)
	if d.CodeSnippet != "" {
		sb.WriteString(fmt.Sprintf("%s```%s\n%s\n%s```\n\n", indent, d.Language, d.CodeSnippet, indent))
	}

	d.PrintBody(&sb)
	return sb.String()
}
