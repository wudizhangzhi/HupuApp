package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/tidwall/gjson"
)

func test() {
	type pepper struct {
		Name     string
		HeatUnit int
		Peppers  int
	}
	peppers := []pepper{
		{Name: "Bell Pepper", HeatUnit: 0, Peppers: 0},
		{Name: "Banana Pepper", HeatUnit: 100, Peppers: 1},
		{Name: "Poblano", HeatUnit: 1000, Peppers: 2},
		{Name: "Jalapeño", HeatUnit: 3500, Peppers: 3},
		{Name: "Aleppo", HeatUnit: 10000, Peppers: 4},
		{Name: "Tabasco", HeatUnit: 30000, Peppers: 5},
		{Name: "Malagueta", HeatUnit: 50000, Peppers: 6},
		{Name: "Habanero", HeatUnit: 100000, Peppers: 7},
		{Name: "Red Savina Habanero", HeatUnit: 350000, Peppers: 8},
		{Name: "Dragon’s Breath", HeatUnit: 855000, Peppers: 9},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Pepper ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
{{ "Peppers:" | faint }}	{{ .Peppers }}`,
	}

	searcher := func(input string, index int) bool {
		pepper := peppers[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Spicy Level",
		Items:     peppers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, peppers[i].Name)
}

type MatchTextMsg struct {
	CommentId         string `json:"commentId"`
	PreviousCommentId string `json:"previousCommentId"`
	NickName          string `json:"nickName"`
	Time              string `json:"time"`
	Content           string `json:"content"`
	Style             string `json:"style"`
	Color             string `json:"color"`
}

func main() {
	matchTextMsgs := []MatchTextMsg{}
	data, err := ioutil.ReadFile("messagegame.txt")
	if err != nil {
		panic(err)
	}
	for _, msg := range gjson.GetBytes(data, "result").Array() {
		matchTextMsg := MatchTextMsg{}
		// fmt.Printf("%+v\n", msg.Value())
		byteResult, _ := json.Marshal(msg.Value())
		json.Unmarshal(byteResult, &matchTextMsg)
		matchTextMsgs = append(matchTextMsgs, matchTextMsg)
		// logger.Info.Printf("比赛消息: %v", msg.Value())
		fmt.Printf("%+v\n", matchTextMsg)
		break
	}
}
