package main

import (
	"time"

	"github.com/alecthomas/kong"
	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/api_utils"
	"github.com/wudizhangzhi/HupuApp/gohupu/live"
	"github.com/wudizhangzhi/HupuApp/gohupu/menu"
)

type LiveCmd struct {
	GameType api.GameType `arg:"" name:"gameType" help:"比赛类型(nba/cba)."`
}

func (r *LiveCmd) Run() error {
	matches, _ := api_utils.GetMatchesFromDate(r.GameType)
	if len(matches) == 0 {
		matches, _ = api_utils.GetAnyMatches(api.NBA, 10, false)
	}
	interfaceItems := make([]interface{}, 0)
	for _, item := range matches {
		interfaceItems = append(interfaceItems, item)
	}
	m := menu.Menu{
		Label:     menu.LabelLive,
		Items:     interfaceItems,
		Templates: menu.LiveTemplate,
		Size:      len(interfaceItems),
	}

	idx, err := m.Start()
	if err != nil {
		return err
	}
	match := matches[idx]
	client := live.Client{
		Match: match,
	}
	// // 退出过快可能导致print打印不显示
	time.Sleep(1 * time.Second)
	client.Start()
	return nil
}

type NewsCmd struct {
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Live LiveCmd `cmd:"" help:"比赛直播."`
	News NewsCmd `cmd:"" help:"新闻."`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("Hupu"),
		kong.Description("A command line tool for Hupu."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree:      true,
			Compact:   true,
			Summary:   true,
			FlagsLast: true,
			// NoExpandSubcommands: true,
		}))
	// TODO 读取配置, 注册signal
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
