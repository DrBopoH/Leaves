// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package database source/database/db.go
package database

import (
	"database/sql"

	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const defaultDbName string = "leaves.db"
const defaultDbPath string = "./db/"

const errorMkdir string = "[EE] DB path dir validation failed with error: "
const errorOpenDb string = "[EE] DB open error: "

const usersTableSQL string = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	password_hash TEXT NOT NULL
);`
const messagesTableSQL string = `
CREATE TABLE IF NOT EXISTS messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY(user_id) REFERENCES users(id)
);`

var DB *sql.DB

func InitDB() error {
	var name string = defaultDbName
	var path string = defaultDbPath

	var err error

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("%s%v\n", errorMkdir, err) // DEBUG PRINT
		log.Fatalf("%s%v\n", errorMkdir, err) // DEBUG LOG
	}

	DB, err = sql.Open("sqlite3", path+name)
	if err != nil {
		fmt.Printf("[WW] SQLite %s DB open failed, aborting...\n", path+name) // DEBUG PRINT

		fmt.Printf("%s%v\n", errorOpenDb, err) // DEBUG PRINT
		log.Fatalf("%s%v\n", errorOpenDb, err) // DEBUG LOG

		return err
	}

	_, err = DB.Exec(usersTableSQL)
	if err != nil {
		fmt.Printf("[WW] SQLite %s DB init table usersTable failed, aborting...\n", path+name)

		fmt.Printf("%s%v\n", errorOpenDb, err)
		log.Fatalf("%s%v\n", errorOpenDb, err)
		return err
	}
	_, err = DB.Exec(messagesTableSQL)
	if err != nil {
		fmt.Printf("[WW] SQLite %s DB init table usersTable failed, aborting...\n", path+name)

		fmt.Printf("%s%v\n", errorOpenDb, err)
		log.Fatalf("%s%v\n", errorOpenDb, err)
		return err
	}
	fmt.Printf("[II] SQLite %s DB ready\n", path+name) // DEBUG PRINT

	return err
}
