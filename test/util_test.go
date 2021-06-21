package main

import (
	"fmt"
	"testing"

	"github.com/wudizhangzhi/HupuApp"
)

func TestLuhnResidue(t *testing.T) {
	var testmap = map[string]int{
		"12345":   1,
		"111111":  9,
		"9527":    5,
		"9999999": 3,
	}
	for k, v := range testmap {
		r := HupuApp.LuhnResidue(k)
		if r != v {
			t.Errorf("LuhnResidue错误，%s 期待: %d, 得到: %d", k, v, r)
		}
	}
}

func TestRandomImei(t *testing.T) {
	imei := HupuApp.GetRandomImei(0, "")
	fmt.Println(imei)
}

func TestGetAndroidId(t *testing.T) {
	HupuApp.GetAndroidId()
}

func TestSign(t *testing.T) {
	params := map[string]string{
		"a": "1",
		"b": "2",
	}
	res := HupuApp.GetSortParam(params)
	if res != "147cbdf43c03bfbdb775aa3ceb7bb729" {
		t.Error("sign不相等")
	}
}
