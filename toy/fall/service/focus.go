package service

import (
	"fmt"
	model "github.com/1211ciel/fall/model/focus"
	"github.com/1211ciel/fall/toy/fall/svc"
	"github.com/1211ciel/fall/utils/fmtt"
	"github.com/1211ciel/fall/utils/input"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func FindMyFocus(c *cli.Context) error {
	user, err := currentUser()
	if err != nil {
		return err
	}
	var datas []model.Focus
	err = svc.DB.Model(&model.Focus{}).Where("uid = ?", user.ID).Find(&datas).Error
	if err != nil {
		color.Red("列表为空")
		return nil
	}
	for _, item := range datas {
		fmtt.Prompt("id\t标题\t类型\t规定时间\t当前完成进度\t总完成时间(分钟)\t总完成次数\t描述\n")
		fmtt.Println(item.ID, "\t", item.Name, "\t", item.TypeCodeDesc(), "\t", item.FlagTime, "\t", float64(item.CurrentTime)/float64(item.FlagTime), "\t", item.TotalTime, "\t", item.FinishNum, "\t", item.Desc)
	}
	return nil
}
func AddMyFocus(c *cli.Context) error {
	u, err := currentUser()
	if err != nil {
		return err
	}
	data := makeNewFocus()
	data.Uid = u.ID
	if err = svc.DB.Create(&data).Error; err != nil {
		return err
	}
	color.HiGreen("创建成功")
	return nil
}
func makeNewFocus() *model.Focus {
	var data model.Focus
	data.Name = input.GetString("请输入专注名称")
	data.TypeCode = uint8(input.GetInt64("请选择类型:0倒计时,1正计时"))
	data.FlagTime = uint32(input.GetInt64("请输入每天完成的目标分钟"))
	data.Desc = input.GetContent("请输入表述信息")
	return &data
}
func TestMakeNewFocus(c *cli.Context) error {
	data := makeNewFocus()
	fmt.Printf("%+v\n", data)
	return nil
}
