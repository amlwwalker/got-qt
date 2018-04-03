package main

import (
	"github.com/jroimartin/gocui"
)

type UIControl struct {
	Views map[string]*gocui.View
	gui *gocui.Gui
}