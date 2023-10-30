package response

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToGetAllMedicines(medicines []domain.Medicine) []web.MedicineResponse {
	var results []web.MedicineResponse
	for _, medicine := range medicines {
		medicineResponse := web.MedicineResponse{
			ID:          int(medicine.ID),
			ScheduleID:	 int(medicine.ScheduleID),
			CategoryID:  int(medicine.CategoryID),
			Name:        medicine.Name,
			Amount:      medicine.Amount,
			Details:     medicine.Details,
			BatchNumber: medicine.BatchNumber,
			// Category:    medicine.Category,
		}
		results = append(results, medicineResponse)
	}
	return results
}

func ConvertToGetMedicine(medicine *domain.Medicine) web.MedicineResponse {
	return web.MedicineResponse{
		ID:          int(medicine.ID),
		ScheduleID:	 int(medicine.ScheduleID),
		CategoryID:  int(medicine.CategoryID),
		Name:        medicine.Name,
		Amount:      medicine.Amount,
		Details:     medicine.Details,
		BatchNumber: medicine.BatchNumber,
		// Category:    medicine.Category,
	}
}
