package __Go语言常用标准库

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestAtoi(t *testing.T) {
	s1 := "100"
	atoi, err := strconv.Atoi(s1)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(atoi)
}
func TestHttp(t *testing.T) {
	apiUrl := "http://localhost:12011/hello"
	data := url.Values{}
	data.Set("name", "Calendar")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		log.Println(err.Error())
		return
	}
	u.RawQuery = data.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(body))
}
func TestGet(t *testing.T) {
	resp, err := resty.New().R().
		SetQueryParam("name", "sciel").
		SetAuthToken("123").
		SetHeader("hehe", "这样就很方便").
		Get("http://localhost:12011/hello")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(string(resp.Body()))
}
func TestPost(t *testing.T) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("token", "123").
		SetBody(`{"name":"john"}`).
		SetHeader("Content-Type", "application/json").
		Post("http://localhost:12011/add")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(string(resp.Body()))
}

type Aoo struct {
}

func TestPrint(t *testing.T) {
	s := Aoo{}
	fmt.Printf("%#v", s)
}

var num = 0

func TestTimer(t *testing.T) {
	action()
}
func action() {
	// 捕获恐慌
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到了")
			action()
		}
	}()
	// 创建一个定时器
	ticker := time.NewTicker(time.Second)
	// 执行动作
	for range ticker.C {
		doSomeThing(ticker)
	}
}

func doSomeThing(t *time.Ticker) {
	num++
	// 动态改变下次执行时间
	t.Reset(time.Second * time.Duration(num))
	fmt.Println(num)
	if num > 5 {
		num = 0
		// 模拟恐慌
		panic("panic")
	}
}
