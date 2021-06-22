package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/live"
)

func TestFetchToken(t *testing.T) {
	addresses := api.GetIpAddress()
	fmt.Printf("获取address: %+v\n", addresses)
	address := HupuApp.RandomChoice(addresses)
	fmt.Printf("选择address: %s\n", address)
	client := live.Client{Domain: address}
	token, err := client.FetchToken()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("token: %s\n", token)
}

func TestLive(t *testing.T) {
	// addresses := api.GetIpAddress()
	// fmt.Printf("获取address: %+v\n", addresses)
	// address := HupuApp.RandomChoice(addresses)
	// fmt.Printf("选择address: %s\n", address)
	games, err := api.GetGameToday(api.NBA)
	fmt.Printf("获取games: %+v\n", games)
	if err != nil {
		t.Error(err)
		return
	}
	if len(games) == 0 {
		games, _ = api.GetGameFromDate(api.NBA, time.Now().AddDate(0, 0, -1).Format("20060102"))
	}
	game := games[len(games)-1]
	fmt.Printf("选择game: %+v\n", game)
	client := live.Client{
		Domain: api.Domain,
		Game:   game,
	}
	client.Start()
}
