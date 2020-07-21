package data

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockDBStore struct {
	mock.Mock
}

//ExecMigration create all tables
func (m *MockDBStore) ExecMigration() bool {
	return true
}

//Select represent select records from database
func (m *MockDBStore) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}

//Get return one record from store
func (m *MockDBStore) Get(dest interface{}, query string, args ...interface{}) error {
	return nil
}

//Insert represent insert dbquery with returning inserted id value
func (m *MockDBStore) Insert(query string, args ...interface{}) (id int) {
	return
}

//Update represent update dbquery
func (m *MockDBStore) Update(query string, args map[string]interface{}) (sql.Result, error) {
	return nil, nil
}
