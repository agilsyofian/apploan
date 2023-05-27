package models

import "gorm.io/gorm"

type Config struct {
	ID       int64   `json:"id" gorm:"column:id"`
	Name     string  `json:"name" gorm:"column:name"`
	Desc     string  `json:"desc" gorm:"column:desc"`
	Constant float64 `json:"constant" gorm:"column:constant"`
	gorm.Model
}

func (db *Database) ConfigGetList() ([]Config, error) {
	var result []Config
	err := db.KreditPlus.Find(&result).Error
	return result, err
}
