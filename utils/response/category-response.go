package response

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToGetAllCategories(categories []domain.Category) []web.CategoryResponse {
	var results []web.CategoryResponse
	for _, category := range categories {
		categoryResponse := web.CategoryResponse{
			ID:          int(category.ID),
			Name:        category.Name,
			Details:     category.Details,
		}
		results = append(results, categoryResponse)
	}
	return results
}

func ConvertToGetCategory(category *domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:          int(category.ID),
		Name:        category.Name,
		Details:     category.Details,
	}
}

