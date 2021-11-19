package db

import (
	"database/sql"

	"gorm.io/gorm"
)

// Create insert the value into database
func (db *DB) Create(value interface{}) (tx *DB) {
	return toDB(db.DB.Create(value))
}

// CreateInBatches insert the value in batches into database
func (db *DB) CreateInBatches(value interface{}, batchSize int) (tx *DB) {
	return toDB(db.DB.CreateInBatches(value, batchSize))
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (db *DB) Save(value interface{}) (tx *DB) {
	return toDB(db.DB.Save(value))
}

// First find first record that match given conditions, order by primary key
func (db *DB) First(dest interface{}, conds ...interface{}) (tx *DB) {
	return toDB(db.DB.First(dest, conds))
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (db *DB) Take(dest interface{}, conds ...interface{}) (tx *DB) {
	return toDB(db.DB.Take(dest, conds))
}

// Last find last record that match given conditions, order by primary key
func (db *DB) Last(dest interface{}, conds ...interface{}) (tx *DB) {
	return toDB(db.DB.Last(dest, conds))
}

// Find find records that match given conditions
func (db *DB) Find(dest interface{}, conds ...interface{}) (tx *DB) {
	return toDB(db.DB.Find(dest, conds))
}

// FindInBatches find records in batches
func (db *DB) FindInBatches(dest interface{}, batchSize int, fc func(tx *DB, batch int) error) *DB {
	return toDB(db.DB.FindInBatches(dest, batchSize, func(tx *gorm.DB, batch int) error {
		return fc(toDB(tx), batch)
	}))
}

func (db *DB) FirstOrInit(dest interface{}, conds ...interface{}) (tx *DB) {
	return toDB(db.DB.FirstOrInit(dest, conds...))
}

func (db *DB) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *DB) {
	return toDB(db.DB.FirstOrCreate(dest, conds...))
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *DB) Update(column string, value interface{}) (tx *DB) {
	return toDB(db.DB.Update(column, value))
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (db *DB) Updates(values interface{}) (tx *DB) {
	return toDB(db.DB.Updates(values))
}

func (db *DB) UpdateColumn(column string, value interface{}) (tx *DB) {
	return toDB(db.DB.UpdateColumn(column, value))
}

func (db *DB) UpdateColumns(values interface{}) (tx *DB) {
	return toDB(db.DB.UpdateColumns(values))
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (db *DB) Delete(value interface{}, conds ...interface{}) (tx *DB) {
	return toDB(db.DB.Delete(value, conds))
}

func (db *DB) Count(count *int64) (tx *DB) {
	return toDB(db.DB.Count(count))
}

func (db *DB) Row() *sql.Row {
	return db.DB.Row()
}

func (db *DB) Rows() (*sql.Rows, error) {
	return db.DB.Rows()
}

// Scan scan value to a struct
func (db *DB) Scan(dest interface{}) (tx *DB) {
	return toDB(db.DB.Scan(dest))
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Model(&users).Pluck("age", &ages)
func (db *DB) Pluck(column string, dest interface{}) (tx *DB) {
	return toDB(db.DB.Pluck(column, dest))
}

func (db *DB) ScanRows(rows *sql.Rows, dest interface{}) error {
	return db.DB.ScanRows(rows, dest)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (db *DB) Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) (err error) {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		return fc(toDB(tx))
	}, opts...)
}

// Begin begins a transaction
func (db *DB) Begin(opts ...*sql.TxOptions) *DB {
	return toDB(db.DB.Begin(opts...))
}

// Commit commit a transaction
func (db *DB) Commit() *DB {
	return toDB(db.DB.Commit())
}

// Rollback rollback a transaction
func (db *DB) Rollback() *DB {
	return toDB(db.DB.Rollback())
}

func (db *DB) SavePoint(name string) *DB {
	return toDB(db.DB.SavePoint(name))
}

func (db *DB) RollbackTo(name string) *DB {
	return toDB(db.DB.RollbackTo(name))
}

// Exec execute raw sql
func (db *DB) Exec(sql string, values ...interface{}) (tx *DB) {
	return toDB(db.DB.Exec(sql, values...))
}
