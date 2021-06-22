package main

import (
	"testing"

	"github.com/wudizhangzhi/HupuApp/gohupu/menu"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

func TestBaseMenu(t *testing.T) {
	items := []message.Game{
		message.Game{
			HomeName:  "湖人",
			AwayName:  "篮网",
			HomeScore: "100",
			AwayScore: "105",
			Process:   "已结束",
		},
		message.Game{
			HomeName:  "76人",
			AwayName:  "老鹰",
			HomeScore: "55",
			AwayScore: "80",
			Process:   "已结束",
		},
	}
	interfaceItems := make([]interface{}, 0)
	for _, item := range items {
		interfaceItems = append(interfaceItems, item)
	}
	m := menu.Menu{
		Label: menu.LabelLive,
		Items: interfaceItems,
		Size:  len(interfaceItems),
	}
	m.Start()
}
