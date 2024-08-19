package main

import (
	"github.com/BottomleyIan/notes-tui/formdata"
	"github.com/BottomleyIan/notes-tui/settings"
	"github.com/rivo/tview"
)

func AddFormDate(form *tview.Form, noteSettings settings.Note, note *formdata.Data) {
	form.AddInputField("Date", note.Date, 0, nil, note.SetDate)
}
func AddFormBody(form *tview.Form, noteSettings settings.Note, note *formdata.Data) {

	if noteSettings.HasBody {
		form.AddTextArea("Body", "", 0, 5, 0, note.SetBody)
	}
}
func AddFormTitle(form *tview.Form, noteSettings settings.Note, note *formdata.Data) {
	if noteSettings.HasTitle {
		form.AddInputField("Title", "", 0, nil, note.SetTitle)
	}
}

func AddFormUrl(form *tview.Form, noteSettings settings.Note, note *formdata.Data) {
	if noteSettings.HasUrl {
		form.AddInputField("Url", "", 0, nil, note.SetUrl)
		form.AddInputField("UrlTitle", "", 0, nil, note.SetUrlTitle)
	}
}

func AddFormCodeSnippet(form *tview.Form, noteSettings settings.Note, note *formdata.Data, languageNames []string) {
	if noteSettings.HasCodeSnippet {
		form.AddDropDown("Language", languageNames, 0, note.SetLanguage)
		form.AddTextArea("Code Snippet", "", 0, 15, 0, note.SetCodeSnippet)
	}
}

func AddFormFields(form *tview.Form, noteSettings settings.Note, note *formdata.Data) {
	for _, f := range noteSettings.Fields {
		if f.Type == "select" {
			form.AddDropDown(f.Name, f.Options, 0, func(option string, optionIndex int) {
				note.SetFieldData(f.Name, option)
			})
		} else {
			form.AddInputField(f.Name, f.DefaultValue, 0, nil, func(value string) {
				note.SetFieldData(f.Name, value)
			})
		}
	}
}
