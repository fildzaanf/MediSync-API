package request

import (
	"app/model/domain"
	"app/model/web"
)

func ConvertToMedicalIDRequest(medicalid web.MedicalIDRequest) *domain.MedicalID {
	return &domain.MedicalID{
		Birthdate:        medicalid.Birthdate,
		Gender:           medicalid.Gender,
		BloodType:        medicalid.BloodType,
		Height:           medicalid.Height,
		Weight:           medicalid.Weight,
		MedicalCondition: medicalid.MedicalCondition,
		EmergencyContact: medicalid.EmergencyContact,
	}
}
