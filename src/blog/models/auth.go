package models

import (
	"crypto/sha256"
	"blog/pkg/setting"
	"fmt"
)

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username,password string)bool{
	var auth Auth
	db.Select("id").Where(Auth{
		Username:username,
		Password:getSaltPassword(password),
	}).First(&auth)

	if auth.ID>0 {
		return true
	}

	return false
}

func AddAuth(username,password string)bool  {
	db.Create(&Auth{
		Username:username,
		Password:getSaltPassword(password),
	})
	return true
}

func getSaltPassword(password string)string{
	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)

	h.Reset()
	h.Write([]byte(fmt.Sprintf("%x",bs)))
	h.Write([]byte(setting.AuthSalt))

	return fmt.Sprintf("%x",h.Sum(nil))
}
