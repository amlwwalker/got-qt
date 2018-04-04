package main

import (
	"github.com/amlwwalker/got-qt/logic"
	"github.com/jroimartin/gocui"
)

type UIControl struct {
	Views map[string]*gocui.View
	gui   *gocui.Gui
	logic *logic.LogicInterface
}
