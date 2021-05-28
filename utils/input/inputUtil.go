package input

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// GetContent 读取控制台输入 以#号结束
func GetContent(msg string) string {
	color.HiGreen("%v,以#结束", msg)
	readString, _ := reader.ReadString('#')
	return readString[0:strings.LastIndex(readString, "#")]
}
func GetString(msg string) string {
	color.HiGreen("%v,以回车结束", msg)
	readString, _ := reader.ReadString('\n')
	return readString[0:strings.LastIndex(readString, "\n")]
}
func GetInt64(msg string) int64 {
	color.HiGreen("%v,以回车结束", msg)
	var data int64
	_, err := fmt.Scan(&data)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return data
}
