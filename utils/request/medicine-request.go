package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToMedicineRequest(medicine web.MedicineRequest, scheduleID int) *domain.Medicine {
	return &domain.Medicine{
		ScheduleID:  uint(scheduleID),
		CategoryID:  uint(medicine.CategoryID), 
		Name:        medicine.Name,
		Amount:      medicine.Amount,
		Details:     medicine.Details,
		BatchNumber: medicine.BatchNumber,
	}
}
