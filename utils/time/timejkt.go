package time

import (
	"log"
	"time"

	"github.com/tigorlazuardi/tanggal"
)

func FormatDateToIndonesian(t time.Time) string {
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
