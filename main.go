package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type myClaim struct {
	Username string `json:"Username"`
	Age      int    `json:"Age"`
	jwt.RegisteredClaims
}

func main() {
	//MapClaims
	//RegisteredClaims
	mc := myClaim{
		Username: "小王",
		Age:      18,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ws",   //签发人
			Subject:   "test", //签发主题
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),    //失效时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(-5 * time.Second)), //生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       //签发时间
			ID:        "",
		},
	}

	//加密
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, mc) //加密方法，加密用户名
	fmt.Println("tok:", tok)
	key := []byte("wyi[iojswoosja") //密钥
	m, err := tok.SignedString(key) //生成密文
	if err != nil {
		fmt.Println("加密出错，", err)
	}
	fmt.Println("密文：", m)
	var claims myClaim
	answer, err1 := jwt.ParseWithClaims(m, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err1 != nil {
		fmt.Println("解密出错，err:", err1)
		return
	}
	if answer.Valid {
		fmt.Println("解密成功，claims:", claims)
	} else {
		fmt.Println("无效的 token")
	}
	fmt.Println(answer.Claims.(*myClaim).Username)
	fmt.Println(answer.Claims.(*myClaim).Age)

}
