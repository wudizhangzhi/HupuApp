package main

import (
	"fmt"
	"testing"

	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/live"
)

func TestFetchToken(t *testing.T) {
	client := live.Client{Domain: HupuApp.DefaultLiveWebSocketDomain}
	token, err := client.FetchToken()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("token: %s\n", token)
}
