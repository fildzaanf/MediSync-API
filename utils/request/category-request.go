package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToCategoryRequest(category web.CategoryRequest) *domain.Category {
	return &domain.Category{
		Name:    category.Name,
		Details: category.Details,
	}
}
