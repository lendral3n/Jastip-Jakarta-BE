package handler

import (
	"jastip-jakarta/features/admin"
	"log"
	"time"

	"github.com/tigorlazuardi/tanggal"
)

type AdminResponse struct {
	ID           uint   `json:"admin_id" form:"admin_id"`
	Name         string `json:"name" form:"name"`
	Role         string `json:"role" form:"role"`
	Email        string `json:"email" form:"email"`
	PhoneNumber  int    `json:"phone_number" form:"phone_number"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
	CreatedAt    string `json:"create_account"`
	UpdatedAt    string `json:"last_update"`
}

type AdminResponseOrder struct {
	Name string `json:"name" form:"name"`
}

func AdminToResponse(data *admin.Admin) AdminResponse {
	return AdminResponse{
		ID:           data.ID,
		Name:         data.Name,
		Email:        data.Email,
		PhoneNumber:  data.PhoneNumber,
		PhotoProfile: data.PhotoProfile,
		Role:         data.Role,
		CreatedAt:    formatDateToIndonesian(data.CreatedAt),
		UpdatedAt:    formatDateToIndonesian(data.UpdatedAt),
	}
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
