package message

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

type WsMsg struct {
	Name string    `json:"name"`
	Args []LiveMsg `json:"args"`
}

type LiveMsg struct {
	Room         string `json:"room"`
	Gid          string `json:"gid"`
	Status       string `json:"status"`
	Pid          int    `json:"pid"`
	RoomLiveType string `json:"room_live_type"`
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
	Uid     int64  `json:"content>uid"`
	Event   string `json:"content>event"`
	EndTime string `json:"content>end_time"`
	T       int    `json:"content>t"`
	// Team    int    `json:"content"`
}
