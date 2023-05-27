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

func (Limit) TableName() string {
	return "limit"
}

func (Kontrak) TableName() string {
	return "kontrak"
}

func (Transaksi) TableName() string {
	return "transaksi"
}
