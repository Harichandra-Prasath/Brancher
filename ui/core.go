package ui

import (
	"github.com/Harichandra-Prasath/Brancher/brancher"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetMainApp(manager *brancher.Manager) *tview.Application {

	app := tview.NewApplication()

	branchList := tview.NewList()
	initList(branchList, manager, app)

	app.SetRoot(branchList, true)
	return app
}

func initList(branchList *tview.List, manager *brancher.Manager, app *tview.Application) {

	branchList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			branchIndex := branchList.GetCurrentItem()
			branchName, _ := branchList.GetItemText(branchIndex)
			manager.BranchCheckout(branchName)
		case 'd':
			branchIndex := branchList.GetCurrentItem()
			branchName, _ := branchList.GetItemText(branchIndex)
			manager.BranchDelete(branchName)
		case 'q':
			app.Stop()
		}

		return event
	})

	for i, branch := range manager.GetLocalBranches() {
		branchList.AddItem(branch, "", rune(49+i), func() {})
	}

	branchList.SetBorder(true).SetTitle("Brancher")

}
