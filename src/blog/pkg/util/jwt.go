package util

import (
	"blog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(username string)(string,error){
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)//过期时间
	claims:=Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)//HASH
	token,err := tokenClaims.SignedString(jwtSecret)//生成签名字符串，再用于获取完整、已签名的token
	return token,err
}

//解析token
func ParseToken(token string)(*Claims,error){
	tokenClaims,err:=jwt.ParseWithClaims(token,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret,nil
	})
	if tokenClaims !=nil{
		if claims,ok:=tokenClaims.Claims.(*Claims);ok&&tokenClaims.Valid{
			//token有效
			return claims,nil
		}
	}
	return nil,err
}