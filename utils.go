package HupuApp

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GetSortParam(params map[string]string) string {
	result := ""
	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if len(result) > 0 {
			result += "&"
		}
		result += strings.Join([]string{key, params[key]}, "=")
	}
	result += HUPU_SALT
	h := md5.New()
	h.Write([]byte(result))
	return hex.EncodeToString(h.Sum(nil))
}

// 随机选择
func RandomChoice(choices []string) string {
	if len(choices) == 0 {
		panic("不可以输入空列表")
	}
	return choices[rand.Intn(len(choices))]
}

func RandomChoiceAny(choices []interface{}) interface{} {
	if len(choices) == 0 {
		panic("不可以输入空列表")
	}
	return choices[rand.Intn(len(choices))]
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
	for i, d := range ReverseStr(digits) {
		d := int(d - '0')
		val := d * (1 + i%2)
		div := val / 10
		mod := val % 10
		result += div
		result += mod
	}
	return result % 10
}

func GetRandomImei(n int, filename string) string {
	if n == 0 {
		n = 15
	}
	return GetImei(n, getRandomTac(filename))
}

// IMEI就是移动设备国际身份码，我们知道正常的手机串码IMEI码是15位数字，
//     由TAC（6位，型号核准号码）、FAC（2位，最后装配号）、SNR（6位，厂商自行分配的串号）和SP（1位，校验位）。
//     tac数据库: https://www.kaggle.com/sedthh/typeallocationtable/data
// part = ''.join(str(random.randrange(0, 9)) for _ in range(N - 1))
// if tac:
//         part = tac + part[len(tac):]
// res = luhn_residue('{}{}'.format(part, 0))
// return '{}{}'.format(part, -res % 10)
func GetImei(n int, tac string) string {
	// part = ''.join(str(random.randrange(0, 9)) for _ in range(N - 1))
	part := ""
	for i := 0; i < n-1; i++ {
		part += fmt.Sprint(rand.Intn(9))
	}
	if tac != "" {
		part = tac + part[len(tac):]
	}
	res := LuhnResidue(part + "0")
	sp := 10 - res%10
	if sp >= 10 {
		sp = sp - 10
	}
	return part + fmt.Sprint(sp)
}

func getRandomTac(filename string) string {
	if filename == "" {
		filename = DefaultTac
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return RandomChoice(TAC_LIST)
	}
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	tacList := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tacList = append(tacList, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return RandomChoice(tacList)
}

// 获取当前时间的millsecond
func GetTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func Str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) - 48) << uint8(i)
	}
	return
}

func GetAndroidId() string {
	result := "0"
	for i := 0; i < 63; i++ {
		result += RandomChoice([]string{"0", "1"})
	}
	base2 := Str2DEC(result)
	return strconv.FormatInt(int64(base2), 16)
}

func InterfaceToStr(i interface{}) string {
	result := ""
	switch i.(type) {

	case string:
		result = fmt.Sprint(i.(string))
	case int:
		result = fmt.Sprint(i.(int))
	case float64:
		result = fmt.Sprint(i.(float64))
	}
	return result
}

func GetTimeFromString(ds string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", ds)
	if err != nil {
		return t, err
	}
	return t, nil
}
