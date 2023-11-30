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
	bbsView := NewBBSView(region)
	bbsView.Display()
}
