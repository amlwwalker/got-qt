package main

import (
	"github.com/amlwwalker/got-qt/logic"
)

func main() {
	cui := &UIControl{}
	//pointer to logic, so needs starting update
	cui.logic = &logic.LogicInterface{}
	cui.logic.ConfigureLogic()
	cui.Init() //init the UI

} // end of main
