package routers

import (
	_ "blog-service/docs"
	"blog-service/global"
	"blog-service/internal/middleware"
	v1 "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// GET  读取和检索
// POST 新增和新建动作
// PUT 更新动作，用于更新一个完整的资源，要求为幂等
// PATCH 更新动作，用于更新某一个资源，的一个组成部分，也就是说，当只需要更新该资源的某一项时，应该使用PATCH而不是PUT， 可以不幂等
// DELETE 删除动作
// HEAD 请求头部信息，用于校验。？？
//

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := NewUpload()

	// 文件上传路由
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	// swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
