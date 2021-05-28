package main

import (
	"fmt"
	"github.com/1211ciel/fall/toy/fall/config"
	"github.com/1211ciel/fall/toy/fall/service"
	"github.com/1211ciel/fall/toy/fall/svc"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	commands = []*cli.Command{

		// 用户管理
		{
			Name:    "user",
			Usage:   "用户管理",
			Aliases: []string{"u"},
			Subcommands: []*cli.Command{
				{
					Name:    "register",
					Usage:   "注册用户",
					Aliases: []string{"r"},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "uname",
							Usage:    "用户名",
							Required: true,
							Aliases:  []string{"u"},
						},
						&cli.StringFlag{
							Name:     "pwd",
							Usage:    "密码",
							Required: true,
							Aliases:  []string{"p"},
						},
					},
					Action: service.Register,
				},
				{
					Name:    "login",
					Usage:   "用户登录",
					Aliases: []string{"l"},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "uname",
							Usage:    "用户名",
							Required: true,
							Aliases:  []string{"u"},
						},
						&cli.StringFlag{
							Name:     "pwd",
							Usage:    "密码",
							Required: true,
							Aliases:  []string{"p"},
						},
					},
					Action: service.Login,
				},
				{
					Name:    "out",
					Usage:   "退出",
					Aliases: []string{"o"},
					Action:  service.Logout,
				},
			},
		},
		// 记事本
		{
			Name:    "note",
			Usage:   "记事本",
			Aliases: []string{"n"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "desc",
					Usage: "说说当下的心情吧",
				},
			},
			Action: service.MyLife,
			Subcommands: []*cli.Command{
				{
					Name:   "time",
					Usage:  "show time",
					Action: service.ShowTime,
				},
			},
		},
		// 专注
		{
			Name:    "focus",
			Usage:   "专注 工作 学习",
			Aliases: []string{"f"},
			Subcommands: []*cli.Command{
				{
					Name:    "add",
					Usage:   "添加专注",
					Aliases: []string{"a"},
					Action:  service.AddMyFocus,
				},
				{
					Name:    "list",
					Usage:   "查看我的专注",
					Aliases: []string{"l"},
					Action:  service.FindMyFocus,
				},
			},
		},
		// 语录
		{
			Name:    "motto",
			Usage:   "语录",
			Aliases: []string{"m"},
		},
		// 测试
		{
			Name:    "test",
			Usage:   "测试",
			Aliases: []string{"t"},
			Subcommands: []*cli.Command{
				{
					Name:    "makeNewFocus",
					Usage:   "创建Focus",
					Aliases: []string{"nf"},
					Action:  service.TestMakeNewFocus,
				},
			},
		},
	}
)

func main() {
	config.NewViper("C:\\Users\\admin\\Desktop\\learn\\fall\\toy\\fall\\config\\config.yaml")
	svc.InitSvc()
	app := cli.NewApp()
	app.Usage = "fall "
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
	}
}
