package html

import (
	"github.com/1211ciel/fall/utils/input"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"strings"
)

func CodeToHtml(*cli.Context) error {
	data := input.GetContent("请输入文本")
	data = strings.ReplaceAll(data, "<", "&lt")
	data = strings.ReplaceAll(data, ">", "&gt")
	color.Green("--------主人,我完成啦--------")
	color.Green(data)
	return nil
}
