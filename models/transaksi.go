package models

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type JenisTransaksi string

const (
	Debit  JenisTransaksi = "debit"
	Kredit JenisTransaksi = "kredit"
)

func (e *JenisTransaksi) Scan(value interface{}) error {
	*e = JenisTransaksi(value.([]byte))
	return nil
}

func (e JenisTransaksi) Value() (driver.Value, error) {
	return string(e), nil
}

type Transaksi struct {
	ID        uuid.UUID      `json:"id" gorm:"column:id"`
	KontrakNo uuid.UUID      `json:"kontrak_no" gorm:"column:kontrak_no"`
	Jenis     JenisTransaksi `json:"jenis" gorm:"column:jenis" sql:"type:ENUM('debit', 'kredit')"`
	Jml       float64        `json:"jml" gorm:"column:jml"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
}

func (db *Database) TransaksiCreate(data Transaksi) (*Transaksi, error) {
	var payload Transaksi = data
	err := db.KreditPlus.Create(&payload).Error
	return &payload, err
}

func (db *Database) TransaksiGetByKontrak(id uuid.UUID) ([]*Transaksi, error) {
	var result []*Transaksi
	err := db.KreditPlus.Where("kontrak_no = ?", id).Find(&result).Error
	return result, err
}

func (db *Database) TransaksiGetByID(id uuid.UUID) (*Transaksi, error) {
	result := &Transaksi{
		ID: id,
	}
	err := db.KreditPlus.Find(&result).Error
	return result, err
}
