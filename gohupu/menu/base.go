package menu

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/manifoldco/promptui"
)

type Menu struct {
	Label     string
	Items     []interface{}
	Templates *promptui.SelectTemplates
	Size      int
}

func (m *Menu) Listen() {
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if char == rune(keyboard.KeyEnter) {
			fmt.Println("回车!")
		}
	}
}

func (m *Menu) Start() (int, error) {
	var idx int
	prompt := promptui.Select{
		Label:     m.Label,
		Items:     m.Items,
		Templates: m.Templates,
		Size:      m.Size,
	}
	idx, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return idx, err
	}
	// fmt.Printf("You choose number %d: %+v\n", i+1, m.Items[i])
	return idx, nil
}
