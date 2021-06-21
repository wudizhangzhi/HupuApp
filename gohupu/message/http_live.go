package message

type MsgGame struct {
	Result struct {
		DayGames []DayGame `json:"games"`
	} `json:"result"`
}

type DayGame struct {
	DateBlock string `json:"date_block"`
	Day       string `json:"day"`
	Games     []Game `json:"data"`
}

// 比赛
type Game struct {
	HomeName  string `json:"home_name"`
	AwayName  string `json:"away_name"`
	HomeScore string `json:"home_score"`
	AwayScore string `json:"away_score"`
	Process   string `json:"process"`
	Gid       string `json:"gid"`
}
