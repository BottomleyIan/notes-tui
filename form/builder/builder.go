package builder

import (
	"github.com/BottomleyIan/notes-tui/formdata"
	"github.com/BottomleyIan/notes-tui/settings"
	"github.com/rivo/tview"
)

func AddUrl(form *tview.Form, noteSettings settings.Note, note *formdata.Data) {
	if noteSettings.HasUrl {
		form.AddInputField("Url", "", 0, nil, note.SetUrl)
		form.AddInputField("UrlTitle", "", 0, nil, note.SetUrlTitle)
	}
}
