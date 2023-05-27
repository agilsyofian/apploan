package models

import (
	"database/sql/driver"
	"time"

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
	AdminFee   float64       `json:"admin_fee" gorm:"column:admin_fee"`
	JmlCicilan float64       `json:"jml_cicilan" gorm:"column:jml_cicilan"`
	JmlBunga   float64       `json:"jml_bunga" gorm:"column:jml_bunga"`
	NamaAsset  string        `json:"nama_asset" gorm:"column:nama_asset" binding:"required"`
	Tenor      int64         `json:"tenor" gorm:"column:tenor;" binding:"required"`
	Status     StatusKontrak `json:"status" gorm:"column:status" sql:"type:ENUM('inpg', 'done','cancel','fail')"`
	gorm.Model
}

type KontrakResponse struct {
	Kontrak Kontrak   `json:"kontrak"`
	Tagihan []Tagihan `json:"tagihan"`
}

func (db *Database) KontrakCreate(data Kontrak) (*KontrakResponse, error) {
	tx := db.KreditPlus.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	var kontrak Kontrak = data
	if err := tx.Create(&kontrak).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tagihan := kontrak.Tagihan()
	if err := tx.Create(&tagihan).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	result := KontrakResponse{
		Kontrak: kontrak,
		Tagihan: tagihan,
	}

	tx.Commit()

	return &result, nil
}

func (k *Kontrak) Tagihan() []Tagihan {
	var tagihan []Tagihan
	curTime := time.Now()
	for i := 0; i < int(k.Tenor); i++ {
		idTagihan, _ := uuid.NewRandom()
		jtp := curTime.AddDate(0, 1, 0)

		tagihan = append(tagihan, Tagihan{
			ID:        idTagihan,
			KontrakNo: k.No,
			Jtp:       jtp.Format("2006-01-02"),
			Jml:       k.JmlCicilan / float64(k.Tenor),
			Status:    "loan",
			CreatedAt: curTime,
			UpdatedAt: curTime,
		})
	}
	return tagihan
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
