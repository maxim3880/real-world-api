package data

import "database/sql"

//Store represent base interface
type Store interface {
	ExecMigration() bool
	Select(interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
	Insert(string, ...interface{}) int
	Update(string, map[string]interface{}) (sql.Result, error)
	Delete(string, ...interface{}) (sql.Result, error)
}
