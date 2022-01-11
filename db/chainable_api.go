package db

import (
	"regexp"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Model specify the model you would like to run db operations
//    // update all users's name to `hello`
//    db.Model(&User{}).Update("name", "hello")
//    // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//    db.Model(&user).Update("name", "hello")
func (db *DB) Model(value interface{}) (tx *DB) {
	return toDB(db.DB.Model(value))
}

// Clauses Add clauses
func (db *DB) Clauses(conds ...clause.Expression) (tx *DB) {
	return toDB(db.DB.Clauses(conds...))
}

var tableRegexp = regexp.MustCompile(`(?i).+? AS (\w+)\s*(?:$|,)`)

// Table specify the table you would like to run db operations
func (db *DB) Table(name string, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Table(name, args))
}

// Distinct specify distinct fields that you want querying
func (db *DB) Distinct(args ...interface{}) (tx *DB) {
	return toDB(db.DB.Distinct(args))
}

// Select specify fields that you want when querying, creating, updating
func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Select(query, args))
}

// Omit specify fields that you want to ignore when creating, updating and querying
func (db *DB) Omit(columns ...string) (tx *DB) {
	return toDB(db.DB.Omit(columns...))
}

// Where add conditions
func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Where(query, args...))
}

// Not add NOT conditions
func (db *DB) Not(query interface{}, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Not(query, args))
}

// Or add OR conditions
func (db *DB) Or(query interface{}, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Or(query, args))
}

// Joins specify Joins conditions
//     db.Joins("Account").Find(&user)
//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
//     db.Joins("Account", DB.Select("id").Where("user_id = users.id AND name = ?", "someName").Model(&Account{}))
func (db *DB) Joins(query string, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Joins(query, args...))
}

// Group specify the group method on the find
func (db *DB) Group(name string) (tx *DB) {
	return toDB(db.DB.Group(name))
}

// Having specify HAVING conditions for GROUP BY
func (db *DB) Having(query interface{}, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Having(query, args...))
}

// Order specify order when retrieve records from database
//     db.Order("name DESC")
//     db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
func (db *DB) Order(value interface{}) (tx *DB) {
	return toDB(db.DB.Order(value))
}

// Limit specify the number of records to be retrieved
func (db *DB) Limit(limit int) (tx *DB) {
	return toDB(db.DB.Limit(limit))
}

// Offset specify the number of records to skip before starting to return the records
func (db *DB) Offset(offset int) (tx *DB) {
	return toDB(db.DB.Offset(offset))
}

// Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically
//     func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
//         return db.Where("amount > ?", 1000)
//     }
//
//     func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
//         return func (db *gorm.DB) *gorm.DB {
//             return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
//         }
//     }
//
//     db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
func (db *DB) Scopes(funcs ...func(*DB) *DB) (tx *DB) {
	fs := make([]func(*gorm.DB) *gorm.DB, 0, len(funcs))

	for _, v := range funcs {
		fs = append(fs, func(in *gorm.DB) *gorm.DB { return v(toDB(in)).DB })
	}
	return toDB(db.DB.Scopes(fs...))
}

// Preload preload associations with given conditions
//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (db *DB) Preload(query string, args ...interface{}) (tx *DB) {
	return toDB(db.DB.Preload(query, args))
}

func (db *DB) Attrs(attrs ...interface{}) (tx *DB) {
	return toDB(db.DB.Attrs(attrs))
}

func (db *DB) Assign(attrs ...interface{}) (tx *DB) {
	return toDB(db.DB.Assign(attrs))
}

func (db *DB) Unscoped() (tx *DB) {
	return toDB(db.DB.Unscoped())
}

func (db *DB) Raw(sql string, values ...interface{}) (tx *DB) {
	return toDB(db.DB.Raw(sql, values...))
}
