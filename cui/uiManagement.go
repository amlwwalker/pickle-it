package main

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
)

func nextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return errors.New("v is not equal to nil")
	}
	switch v.Name() {
	case "menu":
		_, err := g.SetCurrentView("watching")
		return err
	case "watching":
		_, err := g.SetCurrentView("patches")
		return err
	case "patches":
		_, err := g.SetCurrentView("menu")
		return err
	default:
		_, err := g.SetCurrentView("menu")
		return err
	}
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	var err error
	if v != nil {
		cx, cy := v.Cursor()
		if err = v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err = v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	var err error
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err = v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err = v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
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
func getInput(title string, g *gocui.Gui) (string, error) {
	maxX, maxY := g.Size()
	if v, err := g.SetView("input", 1, (maxY/2)-3, maxX-1, (maxY/2)-1); err != nil {
		if err != gocui.ErrUnknownView {
			return "", err
		}
		v.Editable = true
		v.Title = title + " ([space] to cancel)"
		if _, err := g.SetCurrentView("input"); err != nil {
			return "", err
		}
	}
	return "", nil
}

func delView(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView(v.Name()); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("menu"); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func createLayout(g *gocui.Gui, name, title string, x, y, width, height int, highlight bool) (*gocui.View, error) {
	if v, err := g.SetView(name, x, y, width, height); err != nil {
		if err != gocui.ErrUnknownView {
			return &gocui.View{}, err
		}
	} else {
		v.Title = title
		v.Highlight = highlight
		v.Autoscroll = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		return v, nil
	}
	return nil, nil
}

func updateViewElements(g *gocui.Gui, viewName string, list []string) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(viewName)
		if err != nil {
			// handle error
			return err
		}
		v.Clear()
		for _, j := range list {
			fmt.Fprintln(v, j)
		}
		return nil
	})
	return nil
}
func clearView(g *gocui.Gui, viewName string) error {
	v, err := g.View(viewName)
	if err != nil {
		// handle error
		return err
	}
	v.Clear()
	return nil
}
