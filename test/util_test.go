package main

import (
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
