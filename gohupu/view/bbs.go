package view

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/spider"
)

func onSelected(index int, mainText string, secondaryText string, shortcut rune) {
	logger.Info.Printf("selected: %d %s %s %c", index, mainText, secondaryText, shortcut)
}

func onChanged(index int, mainText string, secondaryText string, shortcut rune) {
	logger.Info.Printf("changed: %d %s %s %c", index, mainText, secondaryText, shortcut)
}

func BBSList(region spider.Region, page int) (content tview.Primitive) {
	// modalShown := false
	pages := tview.NewPages()

	leftListView := tview.NewList()
	leftListView.SetHighlightFullLine(true)
	leftListView.SetBorder(true).SetTitle("BBS")
	leftListView.SetSelectedFunc(onSelected)
	leftListView.SetChangedFunc(onChanged)

	rightListView := tview.NewList()
	rightListView.SetBorder(true).SetTitle("Detail")

	bbsList, _ := spider.GetBBSList(region, page)
	for i, bbs := range bbsList[:10] {
		leftListView.AddItem(bbs.Title, fmt.Sprintf("亮:%d 回复:%d", bbs.LightCnt, bbs.ReplyCnt), '0'+rune(i), nil)
	}
	flex := tview.NewFlex().
		AddItem(leftListView, 0, 1, true).
		// AddItem(tview.NewFlex().
		// 	SetDirection(tview.FlexRow).
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Flexible width"), 0, 1, false).
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Fixed height"), 15, 1, false).
		// 	AddItem(tview.NewBox().SetBorder(true).SetTitle("Flexible height"), 0, 1, false), 0, 1, false).
		AddItem(rightListView, 0, 2, false)
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// if modalShown {
		// 	// nextSlide()
		// 	modalShown = false
		// } else {
		// 	pages.ShowPage("modal")
		// 	modalShown = true
		// }
		return event
	})
	modal := tview.NewModal().
		SetText("Resize the window to see the effect of the flexbox parameters").
		AddButtons([]string{"Ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage("modal")
	})
	pages.AddPage("BBS", flex, true, true).
		AddPage("modal", modal, false, false)
	return pages
}
