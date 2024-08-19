package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

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
	note := formdata.New(noteSettings)
	app.form.SetFieldBackgroundColor(tcell.NewRGBColor(0, 0, 0)).
		SetFieldTextColor(tcell.ColorWhite)
	AddFormDate(app.form, noteSettings, note)
	AddFormTitle(app.form, noteSettings, note)
	AddFormFields(app.form, noteSettings, note)
	AddFormUrl(app.form, noteSettings, note)
	AddFormCodeSnippet(app.form, noteSettings, note, app.settings.LanguageNames())
	AddFormBody(app.form, noteSettings, note)
	app.form.AddButton("Save", func() {
		saveJournalEntry(app, note.String(), note.Date)
		app.pages.SwitchToPage("main")
	}).AddButton("Cancel", func() {

		saveJournalEntry(app, note.String(), note.Date)
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

func saveJournalEntry(app *Application, note string, date string) {
	usr, _ := user.Current()
	dir := filepath.Join(usr.HomeDir, app.folder)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	fileName := filepath.Join(dir, "journals", strings.ReplaceAll(date, "-", "_")+".md")
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
