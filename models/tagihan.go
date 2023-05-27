package models

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
)

type StatusTagihan string

const (
	Loan StatusTagihan = "loan"
	Paid StatusTagihan = "paid"
)

func (e *StatusTagihan) Scan(value interface{}) error {
	*e = StatusTagihan(value.([]byte))
	return nil
}

func (e StatusTagihan) Value() (driver.Value, error) {
	return string(e), nil
}

type Tagihan struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	KontrakNo uuid.UUID `json:"kontrak_no" gorm:"column:kontrak_no"`
	Jtp       string    `json:"jtp" gorm:"column:jtp"`
	Jml       float64   `json:"jml" gorm:"column:jml"`
	Status    string    `json:"status" gorm:"column:status" sql:"type:ENUM('loan', 'paid')"`
	TglPaid   *string   `json:"tgl_paid" gorm:"column:tgl_paid"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (db *Database) TagihanExist(konsumenID uuid.UUID) (float64, error) {
	var Cicilan struct {
		Jml float64
	}
	model := db.KreditPlus.Model(&Konsumen{})
	model.Select("SUM(kontrak.jml_cicilan/kontrak.tenor) as jml").Joins("JOIN kontrak ON konsumen.id = kontrak.konsumen_id AND status = 'inpg' AND konsumen.id = ?", konsumenID)
	err := model.Scan(&Cicilan).Error
	if err != nil {
		return 0, err
	}
	return Cicilan.Jml, nil
}
