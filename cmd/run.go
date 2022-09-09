package main

import (
	"log"
	"time"

	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/live"
	"github.com/wudizhangzhi/HupuApp/gohupu/menu"
)

func main() {
	matches, err := api.GetMatchesFromDate(api.NBA)
	if err != nil {
		log.Fatal(err)
		return
	}
	// TODO 测试用, 今日无比赛
	if len(matches) == 0 {
		matches, _ = api.GetMatchesFromDate(api.NBA, time.Now().AddDate(0, 0, -110).Format("20060102"))
	}
	interfaceItems := make([]interface{}, 0)
	for _, item := range matches {
		interfaceItems = append(interfaceItems, item)
	}
	m := menu.Menu{
		Label:     menu.LabelLive,
		Items:     interfaceItems,
		Templates: menu.LiveTemplate,
		Size:      len(interfaceItems),
	}

	idx, err := m.Start()
	if err != nil {
		log.Fatal(err)
		return
	}
	match := matches[idx]
	client := live.Client{
		Domain: api.Domain,
		Match:  match,
	}
	// // 退出过快可能导致print打印不显示
	time.Sleep(1 * time.Second)
	client.Start()
}
