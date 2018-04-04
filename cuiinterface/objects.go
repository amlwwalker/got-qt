package main

import (
	"github.com/jroimartin/gocui"
	"github.com/amlwwalker/got-qt/logic"
)

type UIControl struct {
	Views map[string]*gocui.View
	gui *gocui.Gui
	logic *logic.LogicInterface
}