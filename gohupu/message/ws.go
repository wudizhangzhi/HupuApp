package message

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

type WsMsg struct {
	Name string    `json:"name"`
	Args []LiveMsg `json:"args"`
}

type ScoreBoard struct {
	HomeTid   string `json:"home_tid"`
	HomeName  string `json:"home_name"`
	HomeScore string `json:"home_score"`
	AwayTid   string `json:"away_tid"`
	AwayName  string `json:"away_name"`
	AwayScore string `json:"away_score"`
	Process   string `json:"process"`
	// Status    int    `json:"status"`
}

func (s ScoreBoard) String() string {
	return fmt.Sprintf("%s %s:%s %s  %s", s.AwayName, s.AwayScore, s.HomeScore, s.HomeName, s.Process)
}

func (s ScoreBoard) ColoredString() string {
	red := color.New(color.FgRed).SprintFunc()
	awayscore, _ := strconv.ParseInt(s.AwayScore, 10, 8)
	homescore, _ := strconv.ParseInt(s.HomeScore, 10, 8)
	if awayscore > homescore {
		return fmt.Sprintf("%s %s:%s %s  %s", s.AwayName, red(s.AwayScore), s.HomeScore, s.HomeName, s.Process)
	} else {
		return fmt.Sprintf("%s %s:%s %s  %s", s.AwayName, s.AwayScore, red(s.HomeScore), s.HomeName, s.Process)
	}
}

type LiveMsg struct {
	Room         string `json:"room"`
	Gid          string `json:"gid"`
	Status       string `json:"status"`
	Pid          int    `json:"pid"`
	RoomLiveType int    `json:"room_live_type"`
	OnLine       string `json:"online"`
	Result       Result `json:"result"`
}

type Result struct {
	Score ScoreBoard `json:"scoreboard"`
	// EventMsgs liveMsgSlice `json:"data"`
	Data []struct {
		EventMsgs []EventMsg `json:"a"`
	} `json:"data"`
}

//  "content": {
// 	"uid": "20564829",
// 	"event": "右侧持球突破！继续转移到弧顶格里芬！果断三分出手！！！！",
// 	"end_time": "葱头",
// 	"t": 1623374555
//   }
type EventMsg struct {
	RowId   int `json:"rowId"`
	Content struct {
		Uid     string `json:"uid"`
		Event   string `json:"event"`
		EndTime string `json:"end_time"`
		T       int    `json:"t"`
		// Team    int    `json:"content"`
	} `json:"content"`
}

func (m EventMsg) String() string {
	return fmt.Sprintf("%s: %s", m.Content.EndTime, m.Content.Event)
}

// 1::
type MsgOne struct{}

// 2::
type MsgTwo struct{}

// '1::/nba_v1'
type MsgNBAStart struct{}

var (
	MsgRespMsgOne          = "2:::"
	MsgRespMsgTwo          = "2::"
	MsgRespMsgTwoConnected = "1::/nba_v1"
	MsgRespMsgNBAStart     = `5::/nba_v1:{"args":[{"roomid":-1,"gid":%s,"pid":%d,"room":"NBA_PLAYBYPLAY_CASINO"}],"name":"join"}`
)
