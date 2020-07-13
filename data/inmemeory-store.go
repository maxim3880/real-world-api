package data

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"database/sql"

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
	rows, err := s.execSelect(query, args...)
	if err != nil {
		return err
	}
	pntVal := reflect.Indirect(reflect.ValueOf(dest))
	modelT := pntVal.Type().Elem()
	for rows.Next() {
		elem := reflect.New(modelT).Elem()
		cols, _ := rows.Columns()
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println(err)
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		for i := 0; i < elem.NumField(); i++ {
			field := strings.Split(modelT.Field(i).Tag.Get("db"), ",")[0]
			if item, ok := m[field]; ok && elem.Field(i).CanSet() && item != nil {

				switch elem.Field(i).Kind() {
				case reflect.String:
					elem.Field(i).SetString(string((item.([]uint8))))
				case reflect.Float32, reflect.Float64:
					elem.Field(i).SetFloat(item.(float64))
				case reflect.Ptr:
					if reflect.ValueOf(item).Kind() == reflect.Bool {
						itemBool := item.(bool)
						elem.Field(i).Set(reflect.ValueOf(&itemBool))
					}
				case reflect.Int:
					intVal, _ := strconv.Atoi(string(item.([]uint8)))
					elem.Field(i).SetInt(int64(intVal))
				case reflect.Struct:
					elem.Field(i).Set(reflect.ValueOf(item))
				default:
					fmt.Println(reflect.TypeOf(item).Kind(), reflect.ValueOf(item))
				}
			}
		}
		newSlice := reflect.Append(pntVal, elem)
		reflect.ValueOf(dest).Elem().Set(newSlice)
	}
	return nil
}

func (s *InMemoryStore) Get(dest interface{}, query string, args ...interface{}) error {
	rows, err := s.execSelect(query, args...)
	if err != nil {
		return err
	}
	pntVal := reflect.Indirect(reflect.ValueOf(dest))
	modelT := pntVal.Type()
	count := 0
	for rows.Next() {
		count++
		elem := reflect.New(modelT).Elem()
		cols, _ := rows.Columns()
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println(err)
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		for i := 0; i < elem.NumField(); i++ {
			field := strings.Split(modelT.Field(i).Tag.Get("db"), ",")[0]
			if item, ok := m[field]; ok && elem.Field(i).CanSet() && item != nil {

				switch elem.Field(i).Kind() {
				case reflect.String:
					elem.Field(i).SetString(string((item.([]uint8))))
				case reflect.Float32, reflect.Float64:
					elem.Field(i).SetFloat(item.(float64))
				case reflect.Ptr:
					if reflect.ValueOf(item).Kind() == reflect.Bool {
						itemBool := item.(bool)
						elem.Field(i).Set(reflect.ValueOf(&itemBool))
					}
				case reflect.Int:
					intVal, _ := strconv.Atoi(string(item.([]uint8)))
					elem.Field(i).SetInt(int64(intVal))
				case reflect.Struct:
					elem.Field(i).Set(reflect.ValueOf(item))
				default:
					fmt.Println(reflect.TypeOf(item).Kind(), reflect.ValueOf(item))
				}
			}
		}

		reflect.ValueOf(dest).Elem().Set(elem)
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (s *InMemoryStore) execSelect(query string, args ...interface{}) (*sql.Rows, error) {
	db, err := sql.Open("ramsql", s.source)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db.Query(query, args...)

}

func (s *InMemoryStore) Insert(query string, args ...interface{}) (id int) {
	db, err := sql.Open("ramsql", s.source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.QueryRow(query+" RETURNING id", args...).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
