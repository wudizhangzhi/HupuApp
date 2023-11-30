package view

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/spider"
)

const help_content = `
Ctrl+C: 退出 Enter: 选择
Ctrl+N: 下一页 Ctrl+P: 上一页
Ctrl+Q: 切换焦点
`

var BBSView *BBSViewT

type BBSViewT struct {
	Region       spider.Region
	Main         *tview.Pages
	BBSListView  *tview.List
	PostView     *tview.TextView
	CommentsView *tview.List
	// 页数
	BBSPage     int
	CommentPage int
	// 显示
	HelpContent string
	BBSList     []spider.BBS
	Comments    []spider.Comment
	PostContent string
	SelectedBBS spider.BBS
}

func NewBBSView(region spider.Region) *BBSViewT {
	pages := tview.NewPages()
	// 左侧
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow)

	leftListView := tview.NewList()
	leftListView.SetHighlightFullLine(true).SetBorder(true)
	leftListView.SetSelectedFunc(onSelected).SetChangedFunc(onChanged)

	help := tview.NewTextView().SetTextColor(tcell.ColorAqua)
	fmt.Fprint(help, help_content)
	leftPanel.AddItem(leftListView, 0, 5, true).AddItem(help, 0, 1, false)

	// 右侧
	postContentView := tview.NewTextView().SetDynamicColors(true).SetChangedFunc(func() {
		app.Draw()
	})
	postContentView.SetBorder(true).SetTitle("Post")
	commentsView := tview.NewList()
	commentsView.SetBorder(true).SetTitle("Comments-第1页")
	postPanel := tview.NewFlex().SetDirection(tview.FlexRow)
	postPanel.AddItem(postContentView, 0, 1, false)
	postPanel.AddItem(commentsView, 0, 1, false)
	// 主界面
	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPanel, 0, 1, true).
		AddItem(postPanel, 0, 2, false)
	// 全局快捷键
	mainFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			app.Stop()
		} else if event.Key() == tcell.KeyCtrlQ {
			// logger.Debug.Printf("Focus: %+v\n", app.GetFocus())
			// 切换焦点
			switch app.GetFocus() {
			case leftListView:
				app.SetFocus(postContentView)
			case postContentView:
				app.SetFocus(commentsView)
			case commentsView:
				app.SetFocus(leftListView)
			default:
				app.SetFocus(leftListView)
			}
		}
		return event
	})
	// 左侧快捷键
	leftPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			BBSView.NextPageBBS()
		} else if event.Key() == tcell.KeyCtrlP {
			BBSView.PrevPageBBS()
		}
		return event
	})
	// 右侧快捷键
	postPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			BBSView.NextPageComments()
		} else if event.Key() == tcell.KeyCtrlP {
			BBSView.PrevPageComments()
		}
		return event
	})

	pages.AddPage("BBS", mainFlex, true, true)

	BBSView = &BBSViewT{
		Region:       region,
		Main:         pages,
		BBSListView:  leftListView,
		PostView:     postContentView,
		CommentsView: commentsView,
		HelpContent:  help_content,
	}
	return BBSView
}

func onSelected(index int, mainText string, secondaryText string, shortcut rune) {
	logger.Info.Printf("selected: %d %s %s %c", index, mainText, secondaryText, shortcut)
	// 选择之后加载帖子内容和评论
	bbs := BBSView.BBSList[index]
	BBSView.PostContent, _ = bbs.GetDetail()
	BBSView.CommentPage = 0 // 重置评论页数
	BBSView.SelectedBBS = bbs
	BBSView.RefreshPost()
	BBSView.NextPageComments()
}

func onChanged(index int, mainText string, secondaryText string, shortcut rune) {
	logger.Info.Printf("changed: %d %s %s %c", index, mainText, secondaryText, shortcut)
}

func (b *BBSViewT) Display() {
	b.NextPageBBS()
	if err := app.SetRoot(b.Main, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// 刷新左侧标题
func (b *BBSViewT) RefreshTitle() {
	b.BBSListView.SetTitle(string(b.Region) + "-" + fmt.Sprintf("第%d页", b.BBSPage))
}

// 刷新评论标题
func (b *BBSViewT) RefreshCommentsTitle() {
	b.CommentsView.SetTitle("Comments-" + fmt.Sprintf("第%d页", b.CommentPage))
}

// 帖子下一页
func (b *BBSViewT) NextPageBBS() {
	b.BBSPage += 1
	BBSView.BBSList, _ = spider.GetBBSList(b.Region, b.BBSPage)
	b.RefreshBBSList()
	b.RefreshTitle()
}

// 帖子上一页
func (b *BBSViewT) PrevPageBBS() {
	b.BBSPage -= 1
	if b.BBSPage < 1 {
		b.BBSPage = 1
	}
	BBSView.BBSList, _ = spider.GetBBSList(b.Region, b.BBSPage)
	b.RefreshBBSList()
	b.RefreshTitle()
}

// 评论下一页
func (b *BBSViewT) NextPageComments() {
	b.CommentPage += 1
	BBSView.Comments, _ = BBSView.SelectedBBS.GetComments(b.CommentPage)
	b.RefreshComments()
	b.RefreshCommentsTitle()
}

// 评论上一页
func (b *BBSViewT) PrevPageComments() {
	b.CommentPage -= 1
	if b.CommentPage < 1 {
		b.CommentPage = 1
	}
	BBSView.Comments, _ = BBSView.SelectedBBS.GetComments(b.CommentPage)
	b.RefreshComments()
	b.RefreshCommentsTitle()
}

// 刷新帖子列表
func (b *BBSViewT) RefreshBBSList() {
	b.BBSListView.Clear()
	for i, item := range b.BBSList {
		b.BBSListView.AddItem(
			fmt.Sprintf("(%d) %s", i+1, item.Title),
			fmt.Sprintf("   亮:%d 回复:%d", item.LightCnt, item.ReplyCnt),
			0,
			nil,
		)
	}
}

// 刷新帖子内容
func (b *BBSViewT) RefreshPost() {
	b.PostView.Clear()
	fmt.Fprint(b.PostView, b.PostContent)
	logger.Debug.Printf("更新post content: %s", b.PostContent)
}

// 刷新评论
func (b *BBSViewT) RefreshComments() {
	b.CommentsView.Clear()
	for i, item := range b.Comments {
		b.CommentsView.AddItem(
			fmt.Sprintf("(%d) %s", i+1, item.Content),
			fmt.Sprintf(
				"   昵称:%s 亮:%d 时间:%s 位置:%s",
				item.Nickname,
				item.LightCnt,
				item.ReplyTime.Format("2006-01-02 15:04:05"),
				item.Location,
			),
			0,
			nil,
		)
	}
}
