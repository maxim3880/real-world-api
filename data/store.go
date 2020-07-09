package data

//Store represent base interface
type Store interface {
	ExecMigration() bool
	Select(interface{}, string)
}
