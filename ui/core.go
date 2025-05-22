package ui

import (
	"github.com/Harichandra-Prasath/Brancher/brancher"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	messageChan = make(chan string)
)

func GetMainApp(manager *brancher.Manager) *tview.Application {

	app := tview.NewApplication()

	mainGrid := tview.NewGrid()

	branchList := tview.NewList()
	initList(branchList, manager, app)

	messageText := tview.NewTextView()
	messageText.SetDynamicColors(true)
	go populateMessageText(messageText, app)

	mainGrid.AddItem(branchList, 0, 0, 4, 10, 0, 0, true)
	mainGrid.AddItem(messageText, 10, 0, 1, 10, 0, 0, true)
	app.SetRoot(mainGrid, true)
	return app
}

func initList(branchList *tview.List, manager *brancher.Manager, app *tview.Application) {

	branchList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			branchIndex := branchList.GetCurrentItem()
			branchName, _ := branchList.GetItemText(branchIndex)
			err := manager.BranchCheckout(branchName)
			if err != nil {
				messageChan <- "[red]" + err.Error()
			} else {

				messageChan <- "[green]" + branchName + "Checked out Successfully"
			}
		case 'd':
			branchIndex := branchList.GetCurrentItem()
			branchName, _ := branchList.GetItemText(branchIndex)
			err := manager.BranchDelete(branchName)
			if err != nil {
				messageChan <- "[red]" + err.Error()
			} else {
				messageChan <- "[green]" + branchName + "Deleted Successfully"
			}
		case 'q':
			app.Stop()
		}

		return event
	})

	for i, branch := range manager.GetLocalBranches() {
		branchList.AddItem(branch, "", rune(49+i), func() {})
	}

	branchList.SetBorder(true).SetBorderColor(tcell.ColorBlue)

}

func populateMessageText(messageText *tview.TextView, app *tview.Application) {

	messageText.SetBorder(true).SetBorderColor(tcell.Color100)
	for {
		select {
		case message := <-messageChan:
			messageText.SetText(message)
			app.Draw()
		}

	}

}
