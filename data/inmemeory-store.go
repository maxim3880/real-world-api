package data

import (
	"fmt"

	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/proullon/ramsql/driver"
)

var inMemorySchema = []string{
	"CREATE TABLE IF NOT EXISTS tags (id BIGSERIAL PRIMARY KEY, name text);",
	"INSERT INTO tags (name) VALUES ('reactjs');",
	"INSERT INTO tags (name) VALUES ('angularjs');",
	"CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY, name text, email text, password text, bio text, image text);",
	"INSERT INTO users (name, email, password, bio, image) VALUES ('Maxim', 'maxim3880@gmail.com', '1234', 'test bio text', '');",
}

var inMemStore Store

//CreateImMemmoryStore represent correct create db store
func CreateImMemmoryStore(source string) Store {
	if source == "" {
		source = "TestRamSqlDataSource"
	}
	inMemStore = &InMemoryStore{source}
	if inMemStore.ExecMigration() {
		return inMemStore
	}
	return nil

}

//InMemoryStore represent main application data store
type InMemoryStore struct {
	source string
}

//ExecMigration create all tables
func (s *InMemoryStore) ExecMigration() bool {
	db, err := sql.Open("ramsql", s.source)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, comm := range inMemorySchema {
		_, err = db.Exec(comm)
		if err != nil {
			fmt.Println(err)
		}
	}
	return true
}

//Select represent select records from database
func (s *InMemoryStore) Select(dest interface{}, query string, args ...interface{}) error {
	db, err := sqlx.Connect("ramsql", s.source)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Select(dest, query, args...)
}

func (s *InMemoryStore) Get(dest interface{}, query string, args ...interface{}) error {
	db, err := sqlx.Connect("ramsql", s.source)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Get(dest, query, args...)
}

func (s *InMemoryStore) Insert(query string, args ...interface{}) (id int) {
	query = query + " RETURNING id"
	db, err := sqlx.Connect("ramsql", s.source)
	if err != nil {
		return
	}
	defer db.Close()
	db.QueryRowx(query, args...).Scan(&id)
	return id
}
