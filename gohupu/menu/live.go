package menu

import "github.com/manifoldco/promptui"

var LabelLive = "比赛直播"

var LiveTemplate = &promptui.SelectTemplates{
	Label:    "{{ . }}",
	Active:   "-> {{ .AwayName }} {{ .AwayScore | red  }}:{{ .HomeScore | red }} {{ .HomeName }} {{ .Process }}",
	Inactive: "   {{ .AwayName }} {{ .AwayScore | red  }}:{{ .HomeScore | red }} {{ .HomeName }} {{ .Process }}",
	Selected: "-> {{ .AwayName }} {{ .AwayScore | red  }}:{{ .HomeScore | red }} {{ .HomeName }} {{ .Process }}",
	Help:     "方向键↑↓控制上下, 回车选择",
}
