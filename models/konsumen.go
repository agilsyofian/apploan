package models

import (
	"time"

	"github.com/agilsyofian/kreditplus/utilitize"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Konsumen struct {
	ID          uuid.UUID      `json:"id" gorm:"column:id"`
	Username    string         `json:"username" gorm:"column:username;" binding:"required"`
	Password    string         `json:"password" gorm:"column:password" binding:"required"`
	NIK         int64          `json:"nik" gorm:"column:nik" binding:"required"`
	FullName    string         `json:"full_name" gorm:"column:full_name" binding:"required"`
	LegalName   string         `json:"legal_name" gorm:"column:legal_name" binding:"required"`
	TempatLahir string         `json:"tempat_lahir" gorm:"column:tempat_lahir" binding:"required"`
	TglLahir    string         `json:"tgl_lahir" gorm:"column:tgl_lahir" binding:"required"`
	Gaji        float64        `json:"gaji" gorm:"column:gaji" binding:"required"`
	FotoKTP     string         `json:"foto_ktp" gorm:"column:foto_ktp" binding:"required"`
	FotoSelfie  string         `json:"foto_selfie" gorm:"column:foto_selfie" binding:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Profile struct {
	Detail ProfileDetail     `json:"detail"`
	Limit  []utilitize.Limit `json:"limit"`
}

type ProfileDetail struct {
	NIK         int64     `json:"nik"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	TempatLahir string    `json:"tempat_lahir"`
	TglLahir    string    `json:"tgl_lahir"`
	Gaji        float64   `json:"gaji"`
	FotoKTP     string    `json:"foto_ktp"`
	FotoSelfie  string    `json:"foto_selfie"`
	CreatedAt   time.Time `json:"created_at"`
}

func (db *Database) AuthKonsumen(username string) (*Konsumen, error) {
	var konsumen *Konsumen
	err := db.KreditPlus.Where("username = ?", username).First(&konsumen).Error
	return konsumen, err
}

func (db *Database) CreateKonsumen(data Konsumen) (*Konsumen, error) {
	var konsumen Konsumen = data
	err := db.KreditPlus.Create(&konsumen).Error
	return &konsumen, err
}

func (db *Database) GetKonsumen(id uuid.UUID) (*Konsumen, error) {
	konsumen := &Konsumen{
		ID: id,
	}
	err := db.KreditPlus.Find(&konsumen).Error
	return konsumen, err
}

func (db *Database) GetListKonsumen(page, limit int) ([]*Konsumen, error) {
	var konsumens []*Konsumen
	err := db.KreditPlus.Offset(page).Limit(limit).Find(&konsumens).Error
	return konsumens, err
}

func (db *Database) UpdateKonsumen(id uuid.UUID, data *Konsumen) (*Konsumen, error) {
	konsumen, err := db.GetKonsumen(id)
	if err != nil {
		return nil, err
	}
	err = db.KreditPlus.Model(&konsumen).Updates(data).Error
	return konsumen, err
}

func (db *Database) SoftDeleteKonsumen(id uuid.UUID) (*Konsumen, error) {
	konsumen := &Konsumen{
		ID: id,
	}
	err := db.KreditPlus.Delete(&konsumen).Error
	return konsumen, err
}

func (db *Database) BuildProfile(k Konsumen) (*Profile, error) {

	var response Profile
	response.Detail = ProfileDetail{
		NIK:         k.NIK,
		FullName:    k.FullName,
		LegalName:   k.LegalName,
		TempatLahir: k.TempatLahir,
		TglLahir:    k.TglLahir,
		Gaji:        k.Gaji,
		FotoKTP:     k.FotoKTP,
		FotoSelfie:  k.FotoSelfie,
		CreatedAt:   k.CreatedAt,
	}

	cicilan, err := db.TagihanExist(k.ID)
	if err != nil {
		return &response, err
	}

	configDB, err := db.ConfigGet(1)
	if err != nil {
		return &response, err
	}
	limit := utilitize.NewFactoryLimit(k.Gaji, configDB.Constant, cicilan)

	response.Limit = limit.BuildLimit()
	return &response, nil
}
