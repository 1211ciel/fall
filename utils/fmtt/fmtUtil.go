package fmtt

import (
	"github.com/1211ciel/fall/utils/numutil"
	"github.com/fatih/color"
)

// Prompt 提示
func Prompt(msg string) {
	color.Yellow(msg)
}
func Success(msg string) {
	color.Green(msg)
}
func Info(msg string) {
	color.White(msg)
}
func Error(msg string) {
	color.Red(msg)
}
func Printf(format string, a ...interface{}) {
	_, err := color.Set(color.Attribute(numutil.RandomInt(30, 38))).Printf(format, a...)
	if err != nil {
		return
	}
}
func Println(a ...interface{}) {
	_, err := color.Set(color.Attribute(numutil.RandomInt(30, 38))).Println(a...)
	if err != nil {
		return
	}
}
