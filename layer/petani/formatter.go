package petani

import (
	"Cangkoel-Final-Project/entity"
	"time"
)

type PetaniFormat struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type PetaniDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func Format(petani entity.Petani) PetaniFormat {
	var formatPetani = PetaniFormat{
		ID:       petani.ID,
		FullName: petani.FullName,
		Email:    petani.Email,
	}

	return formatPetani
}

func FormatDelete(msg string) PetaniDeleteFormat {
	var deleteFormat = PetaniDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
