package message

type ScoreBoard struct {
	HomeTid   int    `json:"home_tid"`
	HomeScore int    `json:"home_score"`
	AwayTid   int    `json:"away_id"`
	AwayScore int    `json:"away_score"`
	Process   string `json:"process"`
	Status    int    `json:"status"`
}

type WsMsg struct {
	Room         string     `json:"room"`
	Gid          string     `json:"gid"`
	Status       string     `json:"status"`
	Pid          string     `json:"pid"`
	RoomLiveType int        `json:"room_live_type"`
	LiveMsgs     []LiveMsg  `json:"result>data"`
	ScoreBoard   ScoreBoard `json:"result>scoreboard"`
}

type LiveMsg struct {
	Uid     int64  `json:"content>uid"`
	Event   string `json:"content>event"`
	EndTime string `json:"content>end_time"`
	Team    int    `json:"content"`
}
