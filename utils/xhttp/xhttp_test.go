package xhttp

import (
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	post, err := Post("http://127.0.0.1:8888/myshadow/push", map[string]interface{}{
		"type":      2,
		"operation": 1,
		"room":      "1",
		"msg":       "hello2222",
		"fromUid":   1,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(post)
}
