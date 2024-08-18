package main

import (
	"github.com/rivo/tview"
)

type Application struct {
	app   *tview.Application
	pages *tview.Pages
	form  *tview.Form
}

func NewApplication() *Application {
	app := tview.NewApplication()
	pages := tview.NewPages()
	form := tview.NewForm()
	return &Application{app: app, pages: pages, form: form}
}

func main() {
	app := NewApplication()
	mainMenu(app)
	quickJournalEntry(app)
	if err := app.app.SetRoot(app.pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func mainMenu(app *Application) {

	list := tview.NewList().
		AddItem("Quick Journal Entry", "Add a quick entry to todays Journal", '1', func() {
			switchToQuickJournalEntry(app)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.app.Stop()
		})
	app.pages.AddPage("main", list, true, true)
}

func switchToQuickJournalEntry(app *Application) {
	app.form.Clear(true)

	app.form.AddInputField("Title", "", 30, nil, nil).
		AddInputField("Body", "", 30, nil, nil).
		AddButton("Save", func() {
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
