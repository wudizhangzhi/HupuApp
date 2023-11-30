package view

import (
	"github.com/rivo/tview"
)

var app *tview.Application

func NewApp() *tview.Application {
	app = tview.NewApplication()
	return app
}
