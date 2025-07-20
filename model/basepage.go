package model

type BasePage struct {
	// 分页操作
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}
