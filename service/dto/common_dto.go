package dto

// 通用ID对应的DTO
type CommonIDDTO struct {
	ID uint `json:"id" uri:"id" form:"id"`
}

// 通用分页DTO
type PagingDTO struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

// 对负数进行判断处理
func (p *PagingDTO) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PagingDTO) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	return p.Limit
}
