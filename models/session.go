package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `json:"id" gorm:"column:id"`
	KonsumenID   uuid.UUID `json:"konsumen_id" gorm:"column:konsumen_id"`
	RefreshToken string    `json:"refresh_token" gorm:"column:refresh_token"`
	UserAgent    string    `json:"user_agent" gorm:"column:user_agent"`
	ClientIP     string    `json:"client_ip" gorm:"column:client_ip"`
	ExpiredAt    time.Time `json:"expired_at" gorm:"column:expired_at"`
	IsBlocked    bool      `json:"isblocked" gorm:"column:isblocked"`
}

func (db *Database) SessionCreate(data Session) (*Session, error) {
	var session Session = data
	err := db.KreditPlus.Create(&session).Error
	return &session, err
}

func (db *Database) SessionGet(id uuid.UUID) (*Session, error) {
	session := &Session{
		ID: id,
	}
	err := db.KreditPlus.Find(&session).Error
	return session, err

}
