package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BottomleyIan/notes-tui/form/builder"
	"github.com/BottomleyIan/notes-tui/formdata"
	"github.com/BottomleyIan/notes-tui/settings"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Application struct {
	app      *tview.Application
	pages    *tview.Pages
	form     *tview.Form
	settings settings.Settings
	folder   string
}

func NewApplication() *Application {
	app := tview.NewApplication()
	pages := tview.NewPages()
	form := tview.NewForm()
	settings := settings.ParseSettingsFile()
	return &Application{app: app, pages: pages, form: form, settings: settings}
}

func main() {
	app := NewApplication()
	flag.StringVar(&app.folder, "folder", "notes/", "folder")
	flag.Parse()

	mainMenu(app)
	quickJournalEntry(app)
	if err := app.app.SetRoot(app.pages, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}

}

func mainMenu(app *Application) {
	list := tview.NewList()
	list.
		SetTitle(fmt.Sprintf("Notes in %s", app.folder)).
		SetBorder(true).
		SetBorderPadding(2, 2, 2, 2)

	for _, note := range app.settings.Notes {
		list.AddItem(note.Title, note.Description, rune(note.Key[0]), func() {
			switchToQuickJournalEntry(app, note)
		})
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.app.Stop()
	})

	app.pages.AddPage("main", list, true, true)
}

func switchToQuickJournalEntry(app *Application, noteSettings settings.Note) {
	app.form.Clear(true)
	note := formdata.New()
	app.form.SetFieldBackgroundColor(tcell.NewRGBColor(0, 0, 0)).
		SetFieldTextColor(tcell.ColorWhite)

	app.form.AddInputField("Title", "", 0, nil, note.SetTitle)
	app.form.AddInputField("Tags", "", 0, nil, note.SetTags)
	app.form.AddDropDown("Language", app.settings.LanguageNames(), 0, note.SetLanguage)
	app.form.AddTextArea("Code Snippet", "", 0, 5, 0, note.SetCodeSnippet)

	app.form.AddTextArea("Body", "", 0, 5, 0, note.SetBody)
	builder.AddUrl(app.form, noteSettings, note)
	app.form.AddButton("Save", func() {
		saveJournalEntry(app, note.String())
		app.pages.SwitchToPage("main")
	}).
		AddButton("Quit", func() {
			app.app.Stop()
		})
	app.pages.SwitchToPage("quickJournalEntry")
}

func quickJournalEntry(app *Application) {
	app.pages.AddPage("quickJournalEntry", app.form, true, false)
}

func saveJournalEntry(app *Application, note string) {
	usr, _ := user.Current()
	dir := filepath.Join(usr.HomeDir, app.folder)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	fileName := filepath.Join(dir, "test.md")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.Create(fileName)

	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(note)

}
