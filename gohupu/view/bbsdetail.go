package view

import (
	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func BBSDetail(nextSlide func()) (title string, content tview.Primitive) {
	// pages := tview.NewPages()
	list := tview.NewList().
		AddItem("A Go package for terminal based UIs", "with a special focus on rich interactive widgets", '1', nextSlide).
		AddItem("Based on github.com/gdamore/tcell", "Like termbox but better (see tcell docs)", '2', nextSlide).
		AddItem("Designed to be simple", `"Hello world" is 5 lines of code`, '3', nextSlide).
		AddItem("Good for data entry", `For charts, use "termui" - for low-level views, use "gocui" - ...`, '4', nextSlide).
		AddItem("Extensive documentation", "Everything is documented, examples in GitHub wiki, demo code for each widget", '5', nextSlide)
	return "Detail", Center(80, 10, list)
}
