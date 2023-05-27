package models

type Tabler interface {
	TableName() string
}

func (Konsumen) TableName() string {
	return "konsumen"
}

func (Session) TableName() string {
	return "session"
}

func (Kontrak) TableName() string {
	return "kontrak"
}

func (Config) TableName() string {
	return "config"
}
