package routers

import (
	"github.com/gin-gonic/gin"
	"blog/pkg/setting"
	"net/http"
	"blog/routers/api/v1"
	"blog/routers/api"
	jwt "blog/middleware"
)

func InitRouter()*gin.Engine  {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	router.POST("/auth", api.GetAuth)

	apiv1:=router.Group("/api/v1")
	apiv1.Use(jwt.JWT())//鉴权中间件
	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags",v1.AddTag)
		apiv1.PUT("/tags/:id",v1.EditTag)
		apiv1.DELETE("/tags/:id",v1.DeleteTag)

		apiv1.GET("/articles",v1.GetArticles)
		apiv1.GET("/article/:id",v1.GetArticle)
		apiv1.POST("/articles",v1.AddArticle)
		apiv1.PUT("/articles/:id",v1.EditArticle)
		apiv1.DELETE("/articles/:id",v1.DeleteArticle)
	}
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	return router
}