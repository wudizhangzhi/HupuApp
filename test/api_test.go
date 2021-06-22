package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

func TestGetIpAddress(t *testing.T) {
	// resp, err := api.StatusInit()
	// if err != nil {
	// 	t.Error(err)
	// }
	// defer resp.Body.Close()
	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Println("response Status : ", resp.Status)
	// fmt.Println("response Headers : ", resp.Header)
	// fmt.Println("response Body : ", string(respBody))
	address := api.GetIpAddress()
	fmt.Printf("IP Address: %+v \n", address)
}

func TestGetGames(t *testing.T) {
	resp, err := api.APIGetGames(api.NBA)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	// fmt.Println("response Body : ", string(respBody))
	msgGame := message.MsgGame{}
	if err := json.Unmarshal(respBody, &msgGame); err != nil {
		t.Error(err)
		return
	}
	if len(msgGame.Result.DayGames) == 0 {
		t.Error("解析MsgGame错误")
		return
	}
	if len(msgGame.Result.DayGames[0].Games) == 0 {
		t.Error("解析DayGames错误")
		return
	}
	fmt.Printf("%+v\n", msgGame.Result.DayGames[0].Games[0])
}
