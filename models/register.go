package models

type Register struct {
	Konsumen Konsumen `json:"konsumen"`
	Limit    []Limit  `json:"limit"`
}

func (db *Database) Register(konsumen Konsumen, limit []Limit) (Register, error) {
	var register Register

	tx := db.KreditPlus.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&konsumen).Error; err != nil {
		tx.Rollback()
		return register, err
	}
	register.Konsumen = konsumen

	if err := tx.Create(&limit).Error; err != nil {
		tx.Rollback()
		return register, err
	}
	register.Limit = limit

	tx.Commit()

	return register, nil
}
