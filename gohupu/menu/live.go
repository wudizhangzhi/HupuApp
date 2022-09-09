package menu

import "github.com/manifoldco/promptui"

var LabelLive = "比赛直播"

var LiveTemplate = &promptui.SelectTemplates{
	Label:    "{{ . }}",
	Active:   "-> {{ .AwayTeamName }} {{ .AwayScore | red  }}:{{ .HomeScore | red }} {{ .HomeTeamName }} {{ .MatchStatusChinese }}",
	Inactive: "   {{ .AwayTeamName }} {{ .AwayScore | red  }}:{{ .HomeScore | red }} {{ .HomeTeamName }} {{ .MatchStatusChinese }}",
	Selected: "-> {{ .AwayTeamName }} {{ .AwayScore | red  }}:{{ .HomeScore | red }} {{ .HomeTeamName }} {{ .MatchStatusChinese }}",
	Help:     "方向键↑↓控制上下, 回车选择",
}
