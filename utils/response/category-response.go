package response

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToGetAllCategories(categories []domain.Category) []web.CategoryResponse {
	var results []web.CategoryResponse
	for _, category := range categories {
		categoryResponse := web.CategoryResponse{
			ID:       int(category.ID),
			Name:     category.Name,
			Details:  category.Details,
			// Medicine: category.Medicine,
		}
		// for _, medicine := range category.Medicine{
		// 	medicineResponse := domain.Medicine{
		// 		UserID: medicine.UserID,
		// 		CategoryID: medicine.CategoryID,
		// 		Name: medicine.Name,
		// 		Amount: medicine.Amount,
		// 		Details: medicine.Details,
		// 		BatchNumber: medicine.BatchNumber,
		// 	}

		// 	categoryResponse.Medicine = append(categoryResponse.Medicine, medicineResponse)
		// }
		results = append(results, categoryResponse)
	}
	return results
}

func ConvertToGetCategory(category *domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:       int(category.ID),
		Name:     category.Name,
		Details:  category.Details,
	}
	// for _, medicine := range category.Medicine{
	// 	medicineResponse := domain.Medicine{
	// 		UserID: medicine.UserID,
	// 		CategoryID: medicine.CategoryID,
	// 		Name: medicine.Name,
	// 		Amount: medicine.Amount,
	// 		Details: medicine.Details,
	// 		BatchNumber: medicine.BatchNumber,
	// 	}

	// 	categoryResponse.Medicine = append(categoryResponse.Medicine, medicineResponse)
	// }
	// return categoryResponse
}


