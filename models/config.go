package models

import "gorm.io/gorm"

type Config struct {
	ID       int64   `json:"id" gorm:"column:id"`
	Name     string  `json:"name" gorm:"column:name"`
	Desc     string  `json:"desc" gorm:"column:desc"`
	Constant float64 `json:"constant" gorm:"column:constant"`
	gorm.Model
}

func (db *Database) ConfigCreate(data Config) (*Config, error) {
	var payload Config = data
	err := db.KreditPlus.Create(&payload).Error
	return &payload, err
}

func (db *Database) ConfigGet(id int64) (*Config, error) {
	result := &Config{
		ID: id,
	}
	err := db.KreditPlus.Find(&result).Error
	return result, err
}
