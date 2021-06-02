package main

import (
	"fmt"
	"github.com/1211ciel/fall/utils/ciel/des"
	"github.com/1211ciel/fall/utils/ciel/html"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	cmds = []*cli.Command{
		{
			Name:    "to-html-code",
			Usage:   "将特殊符号转换成html",
			Aliases: []string{"c"},
			Action:  html.CodeToHtml,
		},
		{
			Name:    "des",
			Usage:   "des 工具",
			Aliases: []string{"d"},
			Subcommands: []*cli.Command{
				{
					Name:   "d",
					Usage:  "解密",
					Action: des.DecryptByDES,
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "ciel"
	app.Usage = "我的小工具"
	app.Commands = cmds
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error", err)
	}
}
