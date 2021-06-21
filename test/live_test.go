package main

import (
	"fmt"
	"testing"

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
	game := games[0]
	fmt.Printf("选择game: %+v\n", game)
	client := live.Client{
		Domain: api.Domain,
		Game:   game,
	}
	client.Start()
}
