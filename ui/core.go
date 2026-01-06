package ui

import (
	"errors"
	"strings"
	"time"

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
	initList(branchList, manager, app, mainGrid)

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
				entry := branch
				if branch == manager.CurrentBranch {
					entry += " [green]*"
				}

				list.AddItem(entry, "", rune(49+i), func() {})
			}

			app.Draw()
		}

	}

}

func initTextView(messageText *tview.TextView, app *tview.Application) {
	messageText.SetDynamicColors(true)
	go populateMessageText(messageText, app)

	messageText.SetBorder(true).SetBorderColor(tcell.ColorGrey).SetBorderAttributes(tcell.AttrBold)
	messageText.SetTitleAlign(tview.AlignLeft).SetTitle("Message")

	pushSuccessMessage("Welcome to Brancher")

}

func initDynamicInput(manager *brancher.Manager, grid *tview.Grid, input *tview.InputField, app *tview.Application, label string, action func(*tview.InputField, *brancher.Manager, ...string), selectedBranch ...string) {
	input.SetLabel(label)
	input.SetFieldBackgroundColor(tcell.ColorBlack)

	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			action(input, manager, selectedBranch...)
			destroyInput(input, grid, app)

		} else if key == tcell.KeyEsc {
			destroyInput(input, grid, app)
		}

	})

	grid.AddItem(input, 9, 0, 1, 10, 0, 0, true)

	app.SetFocus(input)
}

func destroyInput(input *tview.InputField, grid *tview.Grid, app *tview.Application) {
	grid.RemoveItem(input)
	app.SetFocus(grid)
}

func initList(branchList *tview.List, manager *brancher.Manager, app *tview.Application, grid *tview.Grid) {

	go activeDrawList(app, branchList, manager)

	branchList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':

			branchName := getSelectedBranch(branchList)

			// Some pre-processing
			branchName = strings.Split(branchName, " ")[0]

			err := manager.BranchCheckout(branchName)
			if err != nil {
				pushErrorMessage(err.Error())
			} else {
				pushSuccessMessage(branchName + " Checked out successfully")
				drawChan <- struct{}{}
			}
		case 'd':

			branchName := getSelectedBranch(branchList)
			branchName = strings.Split(branchName, " ")[0]

			err := manager.BranchDelete(branchName)
			if err != nil {
				pushErrorMessage(err.Error())
			} else {
				pushSuccessMessage(branchName + " deleted successfully")
				drawChan <- struct{}{}
			}
		case 'p':
			branchName := getSelectedBranch(branchList)
			branchName = strings.Split(branchName, " ")[0]

			err := manager.BranchPull(branchName)
			if err != nil {
				if errors.Is(err, brancher.ALREDY_UPTO_DATE) {
					pushSuccessMessage(branchName + " already upto date")
					drawChan <- struct{}{}
				} else {
					pushErrorMessage(err.Error())
				}
			} else {
				drawChan <- struct{}{}
				pushSuccessMessage(branchName + " pulled to latest commit: " + manager.CurrentCommit)
			}

		case 'n':
			input := tview.NewInputField()
			initDynamicInput(manager,
				grid,
				input,
				app,
				"Enter the Branch Name: ",
				creationOperation)
		case 'r':
			input := tview.NewInputField()
			initDynamicInput(manager,
				grid,
				input,
				app,
				"Enter the New Name to Rename: ",
				renameOperation,
				getSelectedBranch(branchList))
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
