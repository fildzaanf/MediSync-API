package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToMedicineRequest(medicine web.MedicineRequest, userID int) *domain.Medicine {
	return &domain.Medicine{
		UserID:  uint(userID),
		CategoryID:  uint(medicine.CategoryID), 
		Name:        medicine.Name,
		Amount:      medicine.Amount,
		Details:     medicine.Details,
		BatchNumber: medicine.BatchNumber,
	}
}
