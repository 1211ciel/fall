package xjwt

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	NewJwt("myShadow", 9999)
	token, err := GenToken(1)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(token)
}

func TestParseTokenDetail(t *testing.T) {
	NewJwt("ciel", 9999)
	claims, err := ParseTokenDetail("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjEsImV4cCI6MTY1ODk3OTEzMywiaXNzIjoic2hhZG93LWltIiwic3ViIjoidXNlci10b2tlbiJ9.MZZWNckCiDNAdT_g_xchsooZjIzP4nTHkwQRRntwehg")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(claims.Uid)
}
func TestPathToken(t *testing.T) {
	NewJwt("ciel", 9999)
	uid, err := PathToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjEsImV4cCI6MTY1ODk3OTEzMywiaXNzIjoic2hhZG93LWltIiwic3ViIjoidXNlci10b2tlbiJ9.MZZWNckCiDNAdT_g_xchsooZjIzP4nTHkwQRRntwehg")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(uid)
}
