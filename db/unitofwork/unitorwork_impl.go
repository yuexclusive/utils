package unitofwork

import "gorm.io/gorm"

type UnitOfWork struct {
	DB *gorm.DB
}

func (u *UnitOfWork) Begin() {
	u.DB.Begin()
}

func (u *UnitOfWork) Commit() {
	u.DB.Commit()
}

func (u *UnitOfWork) Rollback() {
	u.DB.Rollback()
}

func New(db *gorm.DB) IUnitOfWork {
	db.Begin()
	return &UnitOfWork{DB: db}
}
