package api

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"blog/pkg/err"
	"log"
	"blog/models"
	"blog/pkg/util"
	"net/http"
	"blog/pkg/logging"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context){
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	valid:=validation.Validation{}
	a:=auth{Username:username,Password:password}
	ok,_:=valid.Valid(&a)

	data:=make(map[string]interface{})
	code:=err.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username,password)
		if isExist{
			token,e:=util.GenerateToken(username)
			if e!=nil {
				code=err.ERROR_AUTH_TOKEN
			}else{
				data["token"] = token
				code = err.SUCCESS
			}
		}

	}else {
		for _, err := range valid.Errors {
			logging.Info(err.Key,err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : err.GetMsg(code),
		"data" : data,
	})
}
