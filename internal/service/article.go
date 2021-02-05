package service

type GetArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"required,oneof=0 1"`
}

type GetArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Name     string `form:"name" binding:"required,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Content  string `form:"content"`
	TagID    uint32 `form:"tag_id" binding:"gte=1"`
	CreateBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	Content    string `form:"content"`
	ModifiedBy string `form:"modified_by" binding:"required,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
