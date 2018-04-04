package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"strconv"
)

type fn func(g *gocui.Gui, v *gocui.View) error

var controls = map[string]fn{} //holder for the functions that the UI controls fire

func searchField(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("search", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			updateView(g, "logs", err.Error())
			return nil
		}
		v.Editable = true

		if _, err := g.SetCurrentView("search"); err != nil {
			updateView(g, "logs", err.Error())
			return nil
		}
	}
	return nil
}

func clearView(g *gocui.Gui, view string) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(view)
		if err != nil {
			return err
		}
		v.Clear()
		return nil
	})
}

func updateView(g *gocui.Gui, view, message string) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(view)
		if err != nil {
			return err
		}
		fmt.Fprintln(v, message)
		//scroll the window if necessary
		cursorDown(g, v)
		return nil
	})
	return nil
}
func (c *UIControl) ExternalLogUpdate(message string) {
	updateView(c.gui, "logs", message)
}
func (c *UIControl) login(g *gocui.Gui, v *gocui.View) error {
	updateView(g, "logs", "logging in")
	return nil
}

func (c *UIControl) searchContact(g *gocui.Gui, v *gocui.View) error {
	//we don't need to set the msg view
	//because its going to take input
	var l string
	var err error
	defer delView("search", g, v)
	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
		//no value searched for...
		return nil
	}
	updateView(g, "logs", "searching for "+l+"...")
	c.logic.SearchForMatches(l, func(p float64, indeterminate bool) {
		//if indeterminate is true, it means the logic
		//has no idea how long it is going to take
		//so cant have progressive update the UI
		if indeterminate {
			updateView(c.gui, "logs", "process takes time....")
		} else {
			updateView(c.gui, "logs", "processing: "+strconv.FormatFloat(p, 'f', 6, 64))
		}
		if p == 1.0 {
			//we have knowledge the process was complete, we can do anything with it
			for _, v := range c.logic.People {
				//the people in the list will be as a result of this function
				updateView(c.gui, "logs", "found: "+v.Email)
				updateView(c.gui, "contacts", v.Email)
			}
		}
	})
	return nil
}

func (c *UIControl) Init() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	c.gui = g
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(c.layout)

	if err := c.keybindings(g); err != nil {
		updateView(c.gui, "logs", "cannot bind keys "+err.Error())
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		updateView(c.gui, "logs", "main loop errord "+err.Error())
	}
}

func (c *UIControl) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("controls", 0, 0, 30, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Controls"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		controls["4. Search Contact"] = searchField
		controls["1. login"] = c.login
		for k, _ := range controls {
			fmt.Fprintln(v, k)
		}
	}
	if v, err := g.SetView("contacts", 30, 0, 60, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Contacts"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		if _, err := g.SetCurrentView("controls"); err != nil {
			return err
		}
	}
	if v, err := g.SetView("files", 60, 0, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Files"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "sample.txt")
		fmt.Fprintln(v, "accounts.pdf")
		fmt.Fprintln(v, "design.svg")

		if _, err := g.SetCurrentView("controls"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("logs", 0, 15, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Logs"
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "loading...")
		fmt.Fprintln(v, "loading complete")
		if _, err := g.SetCurrentView("controls"); err != nil {
			return err
		}
	}

	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "controls" {
		updateView(g, "logs", "setting view to contacts")
		_, err := g.SetCurrentView("contacts")
		return err
	} else if v.Name() == "contacts" {
		updateView(g, "logs", "setting view to files")
		_, err := g.SetCurrentView("files")
		return err
	}
	updateView(g, "logs", "setting view to controls")
	_, err := g.SetCurrentView("controls")
	return err
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func getControl(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	for i, _ := range controls {
		if l == i {
			//need a helper functon to know what to do with the view
			controls[i](g, v)
			break
		}
	}
	return nil
}

func getLine(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}
func delMsg(g *gocui.Gui, v *gocui.View) error {
	return delView("msg", g, v)
}
func delView(view string, g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView(view); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("controls"); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (c *UIControl) keybindings(g *gocui.Gui) error {
	//main window bindings
	if err := g.SetKeybinding("controls", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("controls", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("controls", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("controls", gocui.KeyEnter, gocui.ModNone, getControl); err != nil {
		return err
	}
	//side window bindings
	if err := g.SetKeybinding("contacts", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("contacts", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("contacts", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("contacts", gocui.KeyEnter, gocui.ModNone, getLine); err != nil {
		return err
	}
	//msg box binding
	if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, delMsg); err != nil {
		return err
	}
	if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, c.searchContact); err != nil {
		return err
	}

	if err := g.SetKeybinding("files", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("files", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("files", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	//all windows
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}
