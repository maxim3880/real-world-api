package data

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/jmoiron/sqlx"
	//Register postgres in memory driver
	_ "github.com/proullon/ramsql/driver"
)

var inMemoryCurrenDate = time.Now()

const inMemorySQLDateFormat = "2006-01-02T15:04:05.9999-07:00"

var inMemorySchema = []string{
	"CREATE TABLE IF NOT EXISTS tags (id BIGSERIAL PRIMARY KEY, name text);",
	"INSERT INTO tags (name) VALUES ('reactjs');",
	"INSERT INTO tags (name) VALUES ('angularjs');",
	"CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY, name text, email text, password text, bio text, image text);",
	"INSERT INTO users (name, email, password, bio, image) VALUES ('Maxim', 'maxim3880@gmail.com', '1234', 'test bio text', '');",
	"CREATE TABLE IF NOT EXISTS articles ( id BIGSERIAL PRIMARY KEY, slug TEXT, title TEXT NOT NULL, description TEXT NOT NULL, body TEXT NOT NULL, createdAt TIMESTAMP, updatedAt TIMESTAMP, author_id INTEGER, PRIMARY KEY (id)); ",
	"CREATE TABLE IF NOT EXISTS user_favorite_articles ( id BIGSERIAL PRIMARY KEY, article_id INTEGER, user_id INTEGER, PRIMARY KEY (id));",
	"CREATE TABLE IF NOT EXISTS tag_in_articles ( id BIGSERIAL PRIMARY KEY, article_id INTEGER, tag_id INTEGER, PRIMARY KEY (id));",
	"INSERT INTO articles (slug,title,description,body,createdAt, updatedAt, author_id)VALUES ('how-to-train-your-dragon','How to train your dragon','So toothless','It a dragon', now(), now(),1);",
	"INSERT INTO articles (slug,title,description,body,createdAt, updatedAt, author_id)VALUES ('how-to-train-your-dragon-2','How to train your dragon 2','So toothless 2','It a dragon', '" + inMemoryCurrenDate.AddDate(0, 0, 1).Format(inMemorySQLDateFormat) + "', '" + inMemoryCurrenDate.AddDate(0, 0, 1).Format(inMemorySQLDateFormat) + "', 1);",
	"CREATE TABLE IF NOT EXISTS user_follows (id BIGSERIAL PRIMARY KEY, follow_user_id INTEGER, user_id    INTEGER, PRIMARY KEY (id));",
}

//CreateImMemmoryStore represent correct create db store
func CreateImMemmoryStore(source string) Store {
	if source == "" {
		source = "TestRamSqlDataSource"
	}
	inMemStore := &InMemoryStore{source}
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

//Get return one record from store
func (s *InMemoryStore) Get(dest interface{}, query string, args ...interface{}) error {
	db, err := sqlx.Connect("ramsql", s.source)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Get(dest, query, args...)
}

//Insert represent insert dbquery with returning inserted id value
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

//Update represent update dbquery
func (s *InMemoryStore) Update(query string, args map[string]interface{}) (sql.Result, error) {
	db, err := sqlx.Connect("ramsql", s.source)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db.NamedExec(query, args)
}

//Delete represent database delete query
func (s *InMemoryStore) Delete(query string, args ...interface{}) (sql.Result, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db.Exec(query, args...)
}
