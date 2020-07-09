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
	"CREATE TABLE IF NOT EXISTS tags (id integer, name text);",
	"INSERT INTO tags (id, name) VALUES ('1', 'reactjs');",
	"INSERT INTO tags (id, name) VALUES ('2', 'angularjs');",
}

//CreateImMemmoryStore represent correct create db store
func CreateImMemmoryStore() Store {
	store := &InMemoryStore{}
	if store.ExecMigration() {
		return store
	}
	return nil
}

//InMemoryStore represent main application data store
type InMemoryStore struct {
}

//ExecMigration create all tables
func (s *InMemoryStore) ExecMigration() bool {
	db, err := sql.Open("ramsql", "TestLoadUserAddresses")
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, comm := range inMemorySchema {
		_, err = db.Exec(comm)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}

//Select represent select records from database
func (s *InMemoryStore) Select(dest interface{}, query string) {
	pntVal := reflect.Indirect(reflect.ValueOf(dest))
	db, err := sql.Open("ramsql", "TestLoadUserAddresses")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	//modelT := reflect.TypeOf(dest).Elem() //.Type().Elem()
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
}
