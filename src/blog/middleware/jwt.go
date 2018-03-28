package middleware

import (
	"github.com/gin-gonic/gin"
	"blog/pkg/err"
	"blog/pkg/util"
	"time"
	"net/http"
)

//生成鉴权中间件
func JWT()gin.HandlerFunc{
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code=err.SUCCESS
		token:=c.Query("token")
		if token ==""{
			code=err.INVALID_PARAMS
		}else {
			claims,e:=util.ParseToken(token)
			if e!=nil {
				code = err.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix()>claims.ExpiresAt {
				code = err.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != err.SUCCESS {
			//状态未经授权
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"msg":err.GetMsg(code),
				"data":data,
			})
			c.Abort()//终止
			return
		}
		c.Next()
	}
}