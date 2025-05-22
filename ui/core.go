package ui

import (
	"github.com/Harichandra-Prasath/Brancher/brancher"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	messageChan = make(chan string)
	drawChan    = make(chan struct{})
)

func GetMainApp(manager *brancher.Manager) *tview.Application {

	app := tview.NewApplication()

	mainGrid := tview.NewGrid()

	branchList := tview.NewList()
	initList(branchList, manager, app)

	messageText := tview.NewTextView()
	initTextView(messageText, app)

	mainGrid.AddItem(branchList, 0, 0, 4, 10, 0, 0, true)
	mainGrid.AddItem(messageText, 10, 0, 1, 10, 0, 0, true)
	app.SetRoot(mainGrid, true)
	return app
}

func activeDrawList(app *tview.Application, list *tview.List, manager *brancher.Manager) {

	for {

		select {

		case <-drawChan:
			// Syncronise
			if err := manager.SyncLocalBranches(); err != nil {
				pushErrorMessage(err.Error())
			}

			// Clear the Current List
			list.Clear()

			for i, branch := range manager.GetLocalBranches() {
				list.AddItem(branch, "", rune(49+i), func() {})
			}

			app.Draw()
		}

	}

}

func pushSuccessMessage(message string) {
	messageChan <- "[green]" + message
}

func pushErrorMessage(message string) {
	messageChan <- "[red]" + message

}

func initTextView(messageText *tview.TextView, app *tview.Application) {
	messageText.SetDynamicColors(true)
	go populateMessageText(messageText, app)

	messageText.SetBorder(true).SetBorderColor(tcell.ColorGrey).SetBorderAttributes(tcell.AttrBold)

}

func initList(branchList *tview.List, manager *brancher.Manager, app *tview.Application) {

	go activeDrawList(app, branchList, manager)

	branchList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			branchIndex := branchList.GetCurrentItem()
			branchName, _ := branchList.GetItemText(branchIndex)
			err := manager.BranchCheckout(branchName)
			if err != nil {
				pushErrorMessage(err.Error())
			} else {

				pushSuccessMessage(branchName + " Checked out Successfully")
			}
		case 'd':
			branchIndex := branchList.GetCurrentItem()
			branchName, _ := branchList.GetItemText(branchIndex)
			err := manager.BranchDelete(branchName)
			if err != nil {
				pushErrorMessage(err.Error())
			} else {
				pushSuccessMessage(branchName + " Deleted Successfully")
				drawChan <- struct{}{}
			}

		case 'q':
			app.Stop()
		}

		return event
	})

	branchList.SetBorder(true).SetBorderColor(tcell.ColorBlue).SetBorderAttributes(tcell.AttrBold)
	branchList.SetTitleAlign(tview.AlignLeft)
	branchList.SetTitle("Local Branches")

	// Signal the initial Drawing
	drawChan <- struct{}{}

}

func populateMessageText(messageText *tview.TextView, app *tview.Application) {

	for {
		select {
		case message := <-messageChan:
			messageText.SetText(message)
			app.Draw()
		}

	}

}
