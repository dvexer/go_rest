/*Package model use sqlite3 for phones database*/
package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go_rest/logger"
)

var db *sql.DB

func init() {
	localDB, err := sql.Open("sqlite3", "./db.sqlite3")
	logger.LogErrorIfExist(err)

	_, err = localDB.Exec(`create table if not exists phones (
		id integer primary key,
		name varchar(30),
		phone varchar(30)
	)`)
	logger.LogErrorIfExist(err)

	db = localDB

	fmt.Println("Initialized!")
}

//CloseDB close database connection
func CloseDB() {
	err := db.Close()
	logger.LogErrorIfExist(err)
}

//Create row in db
func Create(name, phone string) int64 {
	res, err := db.Exec(`insert into phones (name, phone)
		values (?, ?)`, name, phone) // prepare statement
	logger.LogErrorIfExist(err)      // reconsider handling such kind of situations
	id, err := res.LastInsertId()
	logger.LogErrorIfExist(err)
	return id
}

//Read row by name from db
func Read(name string) (phone string) {
	rows, err := db.Query(`select phone from phones where name=?`, name)
	logger.LogErrorIfExist(err)

	if rows.Next() {
		err := rows.Scan(&phone)
		logger.LogErrorIfExist(err)
	}

	return
}

//ReadAll rows from db
func ReadAll() [][]string {
	rows, err := db.Query(`select name, phone from phones`)
	logger.LogErrorIfExist(err)

	res := make([][]string, 0)

	for rows.Next() {
		c := make([]string, 2)
		_ = rows.Scan(&c[0], &c[1])
		res = append(res, c)
	}

	return res
}

//Delete row by name from db
func Delete(name string) {
	_, err := db.Exec(`delete from phones where name=?`, name)
	logger.LogErrorIfExist(err)
}

//Update row by name in db
func Update(name, newName, newPhone string) {
	_, err := db.Exec(`update phones set name = ?, phone = ? where name = ?`, newName, newPhone, name)
	logger.LogErrorIfExist(err)
}
