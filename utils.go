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

func Sum(digits []int) int {
	r := 0
	for _, d := range digits {
		r += d
	}
	return r
}

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// def luhn_residue(digits):
//     return sum(sum(divmod(int(d) * (1 + i % 2), 10)) for i, d in enumerate(digits[::-1])) % 10
func LuhnResidue(digits string) int {
	var result int
	// sum(divmod(int(d) * (1 + i % 2), 10)) for i, d in enumerate(digits[::-1])
	for i, d := range ReverseStr(digits) {
		val := int(d) * (1 + i%2)
		div := val / 10
		mod := val % 10
		result += div
		result += mod
	}
	return result % 10
}

func GetImei(i int8, filename string) string {
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
