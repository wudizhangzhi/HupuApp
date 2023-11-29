package view

import (
	// "fmt"
	// "strconv"

	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/wudizhangzhi/HupuApp/gohupu/spider"
)

var app = tview.NewApplication()

func Display(region spider.Region) {
	layout := BBSList(region, 0)
	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
