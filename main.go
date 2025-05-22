package main

import (
	"github.com/Harichandra-Prasath/Brancher/brancher"
	"github.com/Harichandra-Prasath/Brancher/ui"
)

func main() {

	manager := brancher.NewManager()
	if err := manager.AcquireLocalRepo(); err != nil {
		panic(err)
	}
	app := ui.GetMainApp(manager)
	if err := app.Run(); err != nil {
		panic(err)

	}
}
