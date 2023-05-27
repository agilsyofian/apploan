package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusKontrak string

const (
	Inpg   StatusKontrak = "inpg"
	Done   StatusKontrak = "done"
	Cancel StatusKontrak = "cancel"
	Fail   StatusKontrak = "fail"
)

func (e *StatusKontrak) Scan(value interface{}) error {
	*e = StatusKontrak(value.([]byte))
	return nil
}

func (e StatusKontrak) Value() (driver.Value, error) {
	return string(e), nil
}

type Kontrak struct {
	No         uuid.UUID     `json:"no" gorm:"column:no;primaryKey"`
	KonsumenID uuid.UUID     `json:"konsumen_id" gorm:"column:konsumen_id;<-:create"`
	Otr        float64       `json:"otr" gorm:"column:otr" binding:"required"`
	AdminFee   float64       `json:"admin_fee" gorm:"column:admin_fee" binding:"required"`
	JmlCicilan float64       `json:"jml_cicilan" gorm:"column:jml_cicilan" binding:"required"`
	JmlBunga   float64       `json:"jml_bunga" gorm:"column:jml_bunga" binding:"required"`
	NamaAsset  string        `json:"nama_asset" gorm:"column:nama_asset" binding:"required"`
	Status     StatusKontrak `json:"status" gorm:"column:status" sql:"type:ENUM('inpg', 'done','cancel','fail')" binding:"required"`
	gorm.Model
}

func (db *Database) KontrakCreate(data Kontrak) (*Kontrak, error) {
	var payload Kontrak = data
	err := db.KreditPlus.Create(&payload).Error
	return &payload, err
}

func (db *Database) KontrakGetByKonsumen(id uuid.UUID) ([]*Kontrak, error) {
	var result []*Kontrak
	err := db.KreditPlus.Where("konsumen_id = ?", id).Find(&result).Error
	return result, err
}

func (db *Database) KontrakGetByID(id uuid.UUID) (*Kontrak, error) {
	result := &Kontrak{
		No: id,
	}
	err := db.KreditPlus.Find(&result).Error
	return result, err
}

func (db *Database) KontrakUpdate(id uuid.UUID, payload *Kontrak) (*Kontrak, error) {
	data, err := db.KontrakGetByID(id)
	if err != nil {
		return nil, err
	}
	err = db.KreditPlus.Model(&data).Updates(payload).Error
	return data, err
}
