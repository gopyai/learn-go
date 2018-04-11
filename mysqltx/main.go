// Ini adalah demo cara query MySQL dengan menggunakan transaction protection
// sehingga aman digunakan secara parallel.
package main

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
	"v/db"
	"v/err"
)

const (
	dbHost = "localhost"
	dbName = "xxx"
	dbUser = "xxx"
	dbPass = "xxx"
	dbPort = 3306
)

var (
	DB = db.MustOpen(dbHost, dbName, dbUser, dbPass, dbPort)

	wg sync.WaitGroup
)

func main() {
	defer DB.Close()
	initDB()

	for {
		// Simulasi menggunakan 3 go routine yang masing-masing menggunakan
		// limit 5.
		wg.Add(3)
		go runDb(5)
		go runDb(5)
		go runDb(5)
		wg.Wait()
	}
}

func initDB() {
	DB.Exec("DELETE FROM request")
	for i := 0; i < 50; i++ {
		tx, e := DB.Begin()
		err.Panic(e)
		r, e := tx.Exec("INSERT INTO request (status) VALUE (?)", "created")
		err.Panic(e)
		reqId, e := r.LastInsertId()
		err.Panic(e)
		_, e = tx.Exec("INSERT INTO step (request_id, status) VALUES (?,?)", reqId, "created")
		err.Panic(e)
		err.Panic(tx.Commit())
	}
}

func runDb(limit int) {
	defer wg.Done()

	rows, e := DB.Query("SELECT id, status FROM request WHERE status='created' LIMIT ?", limit)
	err.Panic(e)
	for rows.Next() {
		var reqId uint64
		var reqStatus string
		err.Panic(rows.Scan(&reqId, &reqStatus))

		wg.Add(1)
		go worker(reqId)
	}
}

func worker(reqId uint64) {
	defer wg.Done()

	DB.TxBlock(
		func(tx *sql.Tx) error {

			fmt.Printf("[%d]: Begin\n", reqId)

			var z bool
			if e := tx.QueryRow(
				`SELECT 1 FROM request
			WHERE
				id=? AND status=?
			FOR UPDATE`,
				reqId, "created").Scan(&z); e != nil {
				return e
			}

			fmt.Printf("[%d]: Select from request for update\n", reqId)
			time.Sleep(time.Second)

			var stepId uint64
			err.Panic(tx.QueryRow(
				`SELECT id FROM step
			WHERE
				request_id=? AND status=?
			FOR UPDATE`,
				reqId, "created").Scan(&stepId))

			fmt.Printf("[%d:%d]: Select from step for update\n", reqId, stepId)
			time.Sleep(time.Second)

			_, e := tx.Exec(
				`UPDATE request SET
				status=?
			WHERE
				id=?`,
				"done", reqId)
			err.Panic(e)

			fmt.Printf("[%d:%d]: Update request status\n", reqId, stepId)
			time.Sleep(time.Second)

			_, e = tx.Exec(
				`UPDATE step SET
				status=?
			WHERE
				id=?`,
				"done", stepId)
			err.Panic(e)

			fmt.Printf("[%d:%d]: Update step status\n", reqId, stepId)
			time.Sleep(time.Second)

			return nil

		})
}
