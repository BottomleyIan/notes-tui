package formdata

import "fmt"

type data struct {
	Title string
	Body  string
}

func New() *data {
	return &data{}
}

func (d *data) SetTitle(title string) {
	d.Title = title
}

func (d *data) SetBody(body string) {
	d.Body = body
}

func (d *data) String() string {
	return fmt.Sprintf("- [[%s]]\n\n%s\n\n", d.Title, d.Body)
}
