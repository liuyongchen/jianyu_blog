package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取一篇文章内容
// @Param id path int false "文章ID"
// @Param name query string false "文章标题" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Article/{id} [get]
// @Router /api/v1/Article/ [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 获取文章列表
// @Param tag_id path int false "标签ID"
// @Param name query string false "文章标题" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Article [get]
// @Router /api/v1/Article/{tag_id} [get]
func (a Article) List(c *gin.Context) {

}

// @Summary "新增文章"
// @Param name body string true "文章标题" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param content body string false "文章内容"
// @Param tag_id body int false "标签ID"
// @Param create_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Article [post]
func (a Article) Create(c *gin.Context) {

}

// @Summary "更新文章"
// @Param id path int true "文章ID"
// @Param name body string false "文章标题"
// /maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param content body string false "文章内容"
// @Param modified_by body string true "修改者" maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Article/{id} [put]
func (a Article) Update(c *gin.Context) {

}

// @Summary "删除文章"
// @Param id path string true "文章ID"
// @Success 200 {string} string “成功”
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/Article/{id} [delete]
func (a Article) Delete(c *gin.Context) {

}
