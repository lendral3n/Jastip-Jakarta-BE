package handler

import (
	"jastip-jakarta/features/user"
	"log"
	"time"

	"github.com/tigorlazuardi/tanggal"
)

type UserResponse struct {
	ID           uint   `json:"user_id" form:"user_id"`
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	PhoneNumber  int    `json:"phone_number" form:"phone_number"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
	CreatedAt    string `json:"create_account"`
	UpdatedAt    string `json:"last_update"`
}

func UserToResponse(data *user.User) UserResponse {
	photoProfileURL, ok := data.PhotoProfile.(string)
	if !ok {
		photoProfileURL = ""
	}

	var result = UserResponse{
		ID:           data.ID,
		Name:         data.Name,
		Email:        data.Email,
		PhoneNumber:  data.PhoneNumber,
		PhotoProfile: photoProfileURL,
		CreatedAt:    formatDateToIndonesian(data.CreatedAt),
		UpdatedAt:    formatDateToIndonesian(data.UpdatedAt),
	}
	return result
}

func formatDateToIndonesian(t time.Time) string {
	tgl, err := tanggal.Papar(t, "Jakarta", tanggal.WIB)
	if err != nil {
		log.Fatal(err)
	}

	// Menggunakan custom formatting
	format := []tanggal.Format{
		tanggal.Hari,
		tanggal.NamaBulan,
		tanggal.Tahun,
	}
	ss := tgl.Format(" ", format)
	return ss
}
