package message

// 直播消息
type LiveMsg struct {
	CommentId         string `json:"commentId"`
	PreviousCommentId string `json:"previousCommentId"`
	NickName          string `json:"nickName"`
	Time              string `json:"time"`
	Content           string `json:"content"`
	Style             string `json:"style"`
	Color             string `json:"color"`
}

// 比赛
type Match struct {
	MatchId            string `json:"matchId"`
	MatchStatus        string `json:"matchStatus"`
	MatchStatusChinese string `json:"matchStatusChinese"`
	HomeScore          int    `json:"homeScore"`
	AwayScore          int    `json:"awayScore"`
	HomeTeamId         string `json:"homeTeamId"`
	AwayTeamId         string `json:"awayTeamId"`
	HomeTeamName       string `json:"homeTeamName"`
	AwayTeamName       string `json:"awayTeamName"`
	BeginTime          int64  `json:"beginTime"`
	ChinaStartTime     int64  `json:"chinaStartTime"`
	MatchTime          string `json:"matchTime"`
}

// 一天的赛程
type Game struct {
	Day       string  `json:"day"`
	DayBlock  string  `json:"dayBlock"`
	MatchList []Match `json:"matchList"`
}

// 比赛日程
type GameSchedule struct {
	ScheduleListStats struct {
		EarliestDate  string `json:"earliestDate"`
		LatestDate    string `json:"latestDate"`
		AnchorMatchId string `json:"anchorMatchId"`
		AnchorGdcId   string `json:"anchorGdcId"`
		CurrentDate   string `json:"currentDate"`
	} `json:"scheduleListStats"`
	GameList []Game `json:"gameList"`
}
