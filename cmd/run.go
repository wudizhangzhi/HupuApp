package main

import (
	"log"
	"time"

	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/live"
	"github.com/wudizhangzhi/HupuApp/gohupu/menu"
)

func main() {
	games, err := api.GetGameToday(api.NBA)
	if err != nil {
		log.Fatal(err)
		return
	}
	// TODO 测试用, 今日无比赛
	// if len(games) == 0 {
	// 	games, _ = api.GetGameFromDate(api.NBA, time.Now().AddDate(0, 0, -1).Format("20060102"))
	// }
	interfaceItems := make([]interface{}, 0)
	for _, item := range games {
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
	game := games[idx]
	client := live.Client{
		Domain: api.Domain,
		Game:   game,
	}
	// 退出过快可能导致print打印不显示
	time.Sleep(1 * time.Second)
	client.Start()
}
