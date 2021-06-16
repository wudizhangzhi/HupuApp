package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

func TestMessageExtract(t *testing.T) {
	data, err := ioutil.ReadFile("livemsg.txt")
	if err != nil {
		t.Error(err)
	}
	var wsmsg message.WsMsg
	err = json.Unmarshal(data, &wsmsg)
	if err != nil {
		t.Error(err)
		return
	}
	// fmt.Printf("score: %+v\n", wsmsg.Args[0].Result.Score)
	// fmt.Printf("Room: %s\n", wsmsg.Args[0].Room)
	if len(wsmsg.Args) <= 0 {
		t.Error("解析错误")
		return
	}
	if len(wsmsg.Args[0].Result.Data[0].EventMsgs) <= 0 {
		t.Error("解析EventMsgs错误")
		return
	}
	// fmt.Println(jsoniter.Get(data, "args", 0, "result", "data", 0).ToString())
}
