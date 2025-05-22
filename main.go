package main

import (
	"github.com/Harichandra-Prasath/Brancher/brancher"
	"github.com/Harichandra-Prasath/Brancher/ui"
)

func main() {

	manager := brancher.NewManager()
	manager.AcquireLocalRepo()
	manager.SyncLocalBranches()
	app := ui.GetMainApp(manager)
	if err := app.Run(); err != nil {
		panic(err)

	}
}
