package live

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/api_utils"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

type Client struct {
	liveActivityKey string        // (内部接口获取的参数)
	Match           message.Match // 比赛(外部接口获取的参数)
	LastCommentId   string        // (内部接口获取，用于比对直播数据最后一次消息的id)
	LastCommentTime string
	Th              *time.Ticker
	ThQueryMatch    *time.Ticker   // 另起一个ticker 更新比赛状态
	InterruptCh     chan os.Signal // 中断信号
}

func (c Client) ColoredScore() string {
	red := color.New(color.FgRed).SprintFunc()
	awayscore, _ := strconv.ParseInt(HupuApp.InterfaceToStr(c.Match.AwayScore), 10, 8)
	homescore, _ := strconv.ParseInt(HupuApp.InterfaceToStr(c.Match.HomeScore), 10, 8)
	if awayscore > homescore {
		return fmt.Sprintf("%s:%d", red(awayscore), homescore)
	} else {
		return fmt.Sprintf("%d:%s", awayscore, red(homescore))
	}
}

func New(match message.Match) *Client {
	return &Client{
		Match: match,
	}
}

func (c *Client) init() {
	liveActivityKey, err := api_utils.GetLiveActivityKey(c.Match.MatchId)
	if err != nil {
		logger.Error.Fatal(err)
	}
	c.liveActivityKey = liveActivityKey

	logger.Info.Printf("liveActivityKey: %s\n", liveActivityKey)
}

// 格式化直播消息
func (c *Client) PrintLiveMsg(msg message.LiveMsg) {
	fmt.Fprintf(color.Output, "%s %s %s %s| %s: %s\n",
		c.Match.AwayTeamName,
		c.ColoredScore(),
		c.Match.HomeTeamName,
		c.Match.MatchStatusChinese,
		msg.NickName,
		msg.Content,
	)
}

// 比赛状态更新
func (c *Client) OnMatchUpdate() {
	match, err := api_utils.GetSingleMatch(c.Match.MatchId)
	if err != nil {
		logger.Error.Fatal(err)
	}
	if match.MatchStatus == "COMPLETED" {
		c.End()
	} else {
		c.Match = match
	}
}

func (c *Client) isNewComment(msg message.LiveMsg) bool {
	// logger.Info.Printf("对比comment: LastCommentTime:%v LastCommentId:%v msg:%+v", c.LastCommentTime, c.LastCommentId, msg)
	if c.LastCommentTime == "" {
		return true
	}
	t, _ := HupuApp.GetTimeFromString(msg.Time)
	t_last, _ := HupuApp.GetTimeFromString(c.LastCommentTime)
	if t.After(t_last) {
		return true
	} else if t.Equal(t_last) {
		if msg.PreviousCommentId == c.LastCommentId {
			return true
		}
	}
	return false
}

func (c *Client) OnLiveMessage() {
	matchTextMsgs, err := api_utils.GetLiveMsgList(c.Match.MatchId, c.liveActivityKey, "")
	if err != nil {
		logger.Error.Fatal(err)
	}
	// 确认比赛状态
	if len(matchTextMsgs) == 0 {
		c.OnMatchUpdate()
	} else {
		for _, msg := range matchTextMsgs {
			if c.isNewComment(msg) {
				c.PrintLiveMsg(msg)
				c.LastCommentId = msg.CommentId
				c.LastCommentTime = msg.Time
			}
		}
	}
}

func (c *Client) End() {
	logger.Info.Println("比赛结束")
	fmt.Println("----- 直播结束了 -----")
	c.InterruptCh <- syscall.SIGQUIT
}

func (c *Client) Start() {
	fmt.Println("----- 正在连接直播间 -----")
	// 初始化
	c.InterruptCh = make(chan os.Signal, 1)
	c.Th = time.NewTicker(HupuApp.LIVE_HEART_BEAT_PERIOD * time.Second)
	c.ThQueryMatch = time.NewTicker(HupuApp.LIVE_HEART_BEAT_PERIOD * 2 * time.Second)

	signal.Notify(c.InterruptCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	c.init()

OutLoop:
	for {
		select {
		case <-c.InterruptCh:
			logger.Info.Println("收到中断信号，退出直播间")
			break OutLoop
		case <-c.Th.C:
			c.OnLiveMessage()
		case <-c.ThQueryMatch.C:
			c.OnMatchUpdate()
		}
	}

	close(c.InterruptCh)
}
