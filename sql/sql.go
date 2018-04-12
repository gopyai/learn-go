// sql
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := "belajar"
	pass := "pass"
	host := "localhost:3306"
	dbname := "belajar"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbname)
	fmt.Println(dsn)

	// Create the database handle, confirm driver is present
	db, e := sql.Open("mysql", dsn)
	panicIf(e)
	defer db.Close()

	_, e = db.Exec("INSERT INTO satu (name, age) VALUES (?,?)", "Ana", 16)
	panicIf(e)
}

func panicIf(e error) {
	if e != nil {
		panic((e))
	}
}
