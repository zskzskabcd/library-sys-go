package api

type Pagination struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size" binding:"max=50"`
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Size
}

func (p *Pagination) Limit() int {
	return p.Size
}
