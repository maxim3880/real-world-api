package data

//Store represent base interface
type Store interface {
	ExecMigration() bool
	Select(interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
	Insert(string, ...interface{}) int
}
