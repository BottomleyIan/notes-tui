package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()
	mainMenu(app, pages)
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func mainMenu(app *tview.Application, pages *tview.Pages) {

	list := tview.NewList().
		AddItem("Option 1", "Option 1", '1', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	pages.AddPage("main", list, true, true)
}
