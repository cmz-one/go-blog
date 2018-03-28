package util

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"blog/pkg/setting"
)

//分页页码获取
func GetPage(c *gin.Context)int   {
	result := 0
	page,_ :=com.StrTo(c.Query("page")).Int()//获取页码
	if page>0 {
		result = (page -1)*setting.PageSize //页码*大小
	}
	return result
}