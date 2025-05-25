package ui

import (
	"github.com/Harichandra-Prasath/Brancher/brancher"
	"github.com/rivo/tview"
)

func creationOperation(input *tview.InputField, manager *brancher.Manager, selectedBranch ...string) {
	branch := input.GetText()
	err := manager.BranchCreate(branch)
	if err != nil {
		pushErrorMessage(err.Error())
	} else {
		pushSuccessMessage(branch + " Created Successfully")
		drawChan <- struct{}{}
	}
}

func renameOperation(input *tview.InputField, manager *brancher.Manager, selectedBranch ...string) {
	newBranchName := input.GetText()
	oldBranchName := selectedBranch[0]

	err := manager.BranchRename(oldBranchName, newBranchName)
	if err != nil {
		pushErrorMessage(err.Error())
	} else {
		pushSuccessMessage(oldBranchName + " Renamed to " + newBranchName + " Successfully")
		drawChan <- struct{}{}
	}

}

func pushSuccessMessage(message string) {
	messageChan <- "[green]" + message
}

func pushErrorMessage(message string) {
	messageChan <- "[red]" + message

}

func getSelectedBranch(branchList *tview.List) string {
	branchIndex := branchList.GetCurrentItem()
	branchName, _ := branchList.GetItemText(branchIndex)
	return branchName
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
