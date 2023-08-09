package dao

import (
	"bulugen-backend-go/service/dto"

	"gorm.io/gorm"
)

// 分页函数的定义
func Paging(p dto.PagingDTO) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}
