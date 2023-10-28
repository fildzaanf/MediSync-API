package response

import (
	"app/model/domain"
	"app/model/web"

)

func ConvertToGetAllMedicalIDs(medicalids []domain.MedicalID) []web.MedicalIDResponse {
	var results []web.MedicalIDResponse
	for _, medicalid := range medicalids {
		medicalIDResponse := web.MedicalIDResponse{
			ID:               int(medicalid.ID),
			UserID:           medicalid.UserID,
			Birthdate:        medicalid.Birthdate,
			Gender:           medicalid.Gender,
			BloodType:        medicalid.BloodType,
			Height:           medicalid.Height,
			Weight:           medicalid.Weight,
			MedicalCondition: medicalid.MedicalCondition,
			EmergencyContact: medicalid.EmergencyContact,
		}
		results = append(results, medicalIDResponse)
	}

	return results
}

func ConvertToGetMedicalID(medicalid *domain.MedicalID) web.MedicalIDResponse {
	return web.MedicalIDResponse{
		ID:               int(medicalid.ID),
		UserID:           medicalid.UserID,
		Birthdate:        medicalid.Birthdate,
		Gender:           medicalid.Gender,
		BloodType:        medicalid.BloodType,
		Height:           medicalid.Height,
		Weight:           medicalid.Weight,
		MedicalCondition: medicalid.MedicalCondition,
		EmergencyContact: medicalid.EmergencyContact,
	}
}
