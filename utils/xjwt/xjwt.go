package xjwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	xJwt    *Jwt
	issuer  = "shadow-im"
	subject = "user-token"

	tokenErr = errors.New("token已过期")
)

type Jwt struct {
	Secret     []byte
	ExpireTime int
	Issuer     string // 签名颁发者
	Subject    string // 签名主题
}

// SetSecret 定期修改 Secret
func SetSecret(secret string) {
	xJwt.Secret = []byte(secret)
}

func NewJwt(secret string, expireTime int) *Jwt {
	j := &Jwt{
		Secret:     []byte(secret),
		ExpireTime: expireTime,
	}
	xJwt = j
	return j
}

type Claims struct {
	Uid uint
	jwt.StandardClaims
}

func GenToken(uid uint) (string, error) {
	if xJwt == nil {
		panic("xJwt is nil")
	}
	c := &Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(xJwt.ExpireTime)).Unix(),
			Issuer:    issuer,
			Subject:   subject,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(xJwt.Secret)
	if err != nil {
		return "", errors.New("签名失败")
	}
	return token, nil
}
func ParseTokenDetail(tokenStr string) (*Claims, error) {
	if xJwt == nil {
		panic("xjwt is nil")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return xJwt.Secret, nil
	})
	if err != nil {
		return nil, tokenErr
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, tokenErr
}
func PathToken(token string) (uint, error) {
	detail, err := ParseTokenDetail(token)
	if err != nil {
		return 0, tokenErr
	}
	return detail.Uid, nil
}
