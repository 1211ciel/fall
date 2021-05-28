package service

import (
	"fmt"
	config2 "github.com/1211ciel/fall/toy/fall/config"
	"github.com/1211ciel/fall/utils/numutil"
	"github.com/fatih/color"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/urfave/cli/v2"
	"time"
)

func MyLife(c *cli.Context) error {
	color.Green("hello ciel")
	desc := c.String("desc")
	if desc != "" {
		color.Blue(fmt.Sprint(desc))
	}
	return nil
}
func ShowTime(c *cli.Context) error {
	config2.C.Time.Status = 1
	if err := config2.V.WriteConfigAs("config.yaml"); err != nil {
		logx.Error(err.Error())
		return err
	}
	var count int
	for {
		count++
		if config2.V.GetInt("time.status") == 0 {
			color.Blue("时钟已停止,共使用了 %v秒", count)
			return nil
		}
		_, _ = color.Set(color.Attribute(numutil.RandomInt(30, 38))).Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
