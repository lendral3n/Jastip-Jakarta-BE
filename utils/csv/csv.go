package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	
)

type UserOrderCSV struct {
	NamaUser            string
	NomorTeleponWhatsapp string
	NomorResiJastip     string
	NomorResi           string
	NomorOrder          string
	KodeWilayah         string
	HargaPerKodeWilayah string
	Berat               string
	NamaBarang          string
	BatchPengiriman     string
}

type CSVGeneratorInterface interface {
	GenerateCSV(filePath string, data []UserOrderCSV) error
}

type CSVGenerator struct {
}

func New() CSVGeneratorInterface {
	return &CSVGenerator{}
}

// GenerateCSV implements CSVGeneratorInterface.
func (c *CSVGenerator) GenerateCSV(filePath string, data []UserOrderCSV) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// writer.Comma = '\t'

	header := []string{
		"Nama User",
		"Nomor Telepon Whatsapp",
		"Nomor Resi Jastip",
		"Nomor Resi",
		"Nomor Order",
		"Kode Wilayah",
		"Harga per Kode Wilayah",
		"Berat",
		"Nama Barang",
		"Batch Pengiriman",
	}
	err = writer.Write(header)
	if err != nil {
		return fmt.Errorf("failed to write headers: %w", err)
	}

	for _, order := range data {
		row := []string{
			order.NamaUser,
			order.NomorTeleponWhatsapp,
			order.NomorResiJastip,
			order.NomorResi,
			order.NomorOrder,
			order.KodeWilayah,
			order.HargaPerKodeWilayah,
			order.Berat,
			order.NamaBarang,
			order.BatchPengiriman,
		}
		err = writer.Write(row)
		if err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}

	return nil
}
