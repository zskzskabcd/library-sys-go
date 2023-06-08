package api

type Pagination struct {
	Page int `form:"page" json:"page" binding:"min=1" default:"1" example:"1"`    // 页码
	Size int `form:"size" json:"size" binding:"max=50" default:"10" example:"10"` // 每页条数
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Size
}

func (p *Pagination) Limit() int {
	return p.Size
}
