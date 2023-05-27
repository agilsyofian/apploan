package models

import (
	"time"

	"github.com/google/uuid"
)

type Limit struct {
	ID         int64     `json:"id" gorm:"column:id"`
	KonsumenID uuid.UUID `json:"konsumen_id" gorm:"column:konsumen_id;<-:create" binding:"required"`
	Tenor      int64     `json:"tenor" gorm:"column:tenor;<-:create" binding:"required"`
	Limit      float64   `json:"limit" gorm:"column:limit" binding:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;<-:create" binding:"required"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" binding:"required"`
}

func (db *Database) LimitCreate(data []Limit) ([]Limit, error) {
	var result []Limit = data
	err := db.KreditPlus.Create(&result).Error
	return result, err
}

func (db *Database) LimitGetByKonsumen(id uuid.UUID) ([]Limit, error) {
	var result []Limit
	err := db.KreditPlus.Where("konsumen_id = ?", id).Find(&result).Error
	return result, err
}

func (db *Database) LimitGetByID(id int64) (*Limit, error) {
	result := &Limit{
		ID: id,
	}
	err := db.KreditPlus.Find(&result).Error
	return result, err
}

func (db *Database) LimitUpdate(id int64, payload *Limit) (*Limit, error) {
	data, err := db.LimitGetByID(id)
	if err != nil {
		return nil, err
	}
	err = db.KreditPlus.Model(&data).Updates(payload).Error
	return data, err
}
