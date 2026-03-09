// source/database/db.go
package database

import (
	"database/sql"

	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const default_db_name string = "leaves.db"
const default_db_path string = "./db/"

const error_mkdir string = "[EE] DB path dir validation failed with error: "
const error_opendb string = "[EE] DB open error: "

const usersTable string = `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	password_hash TEXT NOT NULL
);`
const messagesTable string = `CREATE TABLE IF NOT EXISTS messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(user_id) REFERENCES users(id)
);`



var DB *sql.DB



func InitDB() error {
	var name string = default_db_name
	var path string = default_db_path
	
	var err error
	

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("%s%v\n", error_mkdir, err)		// DEBUG PRINT
		log.Fatalf("%s%v\n", error_mkdir, err)		// DEBUG LOG
	}


	DB, err = sql.Open("sqlite3", path+name)
	if err != nil {
		fmt.Printf("%s%v\n", error_opendb, err)		// DEBUG PRINT
		log.Fatalf("%s%v\n", error_opendb, err)		// DEBUG LOG
	
		fmt.Printf("[WW] SQLite %s DB open failed, aborting...\n", path+name)	// DEBUG PRINT
		
		return err
	}	


	DB.Exec(usersTable)
	DB.Exec(messagesTable)
	fmt.Printf("[II] SQLite %s DB ready\n", path+name)	// DEBUG PRINT

	return err
}