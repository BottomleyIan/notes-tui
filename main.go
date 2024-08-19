package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BottomleyIan/notes-tui/formdata"
	"github.com/rivo/tview"
)

type Application struct {
	app    *tview.Application
	pages  *tview.Pages
	form   *tview.Form
	folder string
}

func NewApplication() *Application {
	app := tview.NewApplication()
	pages := tview.NewPages()
	form := tview.NewForm()
	return &Application{app: app, pages: pages, form: form}
}

func main() {
	app := NewApplication()
	flag.StringVar(&app.folder, "folder", "notes/", "folder")
	flag.Parse()

	mainMenu(app)
	quickJournalEntry(app)
	if err := app.app.SetRoot(app.pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func mainMenu(app *Application) {
	list := tview.NewList()
	list.
		SetTitle(fmt.Sprintf("Notes in %s", app.folder)).
		SetBorder(true).
		SetBorderPadding(2, 2, 2, 2)
	list.AddItem("Quick Journal Entry", "Add a quick entry to todays Journal", '1', func() {
		switchToQuickJournalEntry(app)
	}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.app.Stop()
		})
	app.pages.AddPage("main", list, true, true)
}

func switchToQuickJournalEntry(app *Application) {
	app.form.Clear(true)
	note := formdata.New()
	app.form.AddInputField("Title", "", 30, nil, func(title string) {
		note.SetTitle(title)
	})

	app.form.AddInputField("Body", "", 30, nil, func(body string) {
		note.SetBody(body)
	})

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
