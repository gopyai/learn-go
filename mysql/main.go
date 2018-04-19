package main_test

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "root"
	dbPass = "password"
	dbName = "db_name"
)

func Example() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
	panicIf(err)
	defer db.Close()

	rows, err := db.Query("DESCRIBE any_table")
	panicIf(err)
	for rows.Next() {
		var (
			vField   []byte
			vType    []byte
			vNull    []byte
			vKey     []byte
			vDefault []byte
			vExtra   []byte
		)
		panicIf(rows.Scan(&vField, &vType, &vNull, &vKey, &vDefault, &vExtra))
		fmt.Println("Scan:", string(vField), string(vType), string(vNull), string(vKey), string(vDefault), string(vExtra))
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
