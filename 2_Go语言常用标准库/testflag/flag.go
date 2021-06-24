package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "张三", "姓名")
	var age int64
	flag.Int64Var(&age, "age", 20, "age")
	flag.Parse()
	fmt.Println(*name, age)
}
