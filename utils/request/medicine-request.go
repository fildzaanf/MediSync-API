package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToMedicineRequest(medicine web.MedicineRequest) *domain.Medicine {
	return &domain.Medicine{
		Name:        medicine.Name,
		Amount:      medicine.Amount,
		Details:     medicine.Details,
		BatchNumber: medicine.BatchNumber,
	}
}
