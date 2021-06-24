package __Go语言常用标准库

import (
	"log"
	"os"
	"testing"
)

func TestIoutils(t *testing.T) {
	file, err := os.OpenFile("./xx2.text", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()
	_, err = file.WriteString("hello\n")
}
