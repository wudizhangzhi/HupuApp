package HupuApp

import (
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GetImei(i int8) string {
	return ""
}

func getRandomTac(filename string) string {
	if filename == "" {
		filename = DefaultTac
	}
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ioutil.ReadAll(f)
	return ""
}
