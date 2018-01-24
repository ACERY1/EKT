package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type EKTDB interface {
}

func Query() {
	db, err := sql.Open("sqlite3", "/home/x/s")
	time.Sleep(3 * time.Second)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	rows, err := db.Query("SELECT `id`, `name`, `password` FROM `test` WHERE `name` = ?", "zhouxun")
	for rows.Next() {
		var id int
		var name, password string
		rows.Scan(&id, &name, &password)
		fmt.Println(id, name, password)
	}
}

func Insert() {
	db, err := sql.Open("sqlite3", "/home/x/s")
	time.Sleep(3 * time.Second)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	stmt, err := db.Prepare("INSERT INTO `test`(`name`, `password`) VALUES(?, ?)")
	res, err := stmt.Exec("zhouxun", "helloworld")
	fmt.Println(res.LastInsertId())
}
