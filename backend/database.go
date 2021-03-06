// database related stuff

package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "oatmessenger.db"
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const tokenSeparator = "."

func generateDB(db *sql.DB) {
	db.Exec(`CREATE TABLE "User" (
		"ID"			INTEGER NOT NULL UNIQUE,
		"login"			TEXT NOT NULL UNIQUE COLLATE NOCASE,
		"email" 		TEXT NOT NULL UNIQUE COLLATE NOCASE,
		"full_name"		TEXT NOT NULL DEFAULT "",
		"password_hash"	BLOB NOT NULL,
		"dialogues" TEXT NOT NULL DEFAULT "",
		PRIMARY KEY("ID")
	);`)
	db.Exec(`CREATE TABLE "Message" (
		"ID"			INTEGER NOT NULL UNIQUE,
		"user_id_from"	INTEGER NOT NULL,
		"user_id_for"	INTEGER NOT NULL,
		"text"			TEXT DEFAULT "",
		"attachments"	TEXT DEFAULT "",
		"datetime"		TEXT NOT NULL,
		PRIMARY KEY("ID" AUTOINCREMENT)
	);`)
	db.Exec(`CREATE TABLE "Token" (
		"user_id"	INTEGER NOT NULL,
		"data"		INTEGER NOT NULL UNIQUE,
		PRIMARY KEY("data")
	);`)
}

func dbHasTables(db *sql.DB) bool {
	_, err := db.Query("select * from users;")
	return err == nil
}

func connectToDB(dbName string) *sql.DB {
	var db *sql.DB
	var err error
	db, _ = sql.Open("sqlite3", dbName)
	err = db.Ping()
	panicIfError(err)
	if !dbHasTables(db) {
		generateDB(db)
	}
	return db
}

// func main() {
// 	os.Remove("oatmessenger.db")
// 	db := connectToDB(dbName)
// 	defer db.Close()
// 	err := addUser(db, "Varabe", "83")
// 	panicIfError(err)
// 	user, err := getUserByID(db, 1)
// 	fmt.Println(user)
// 	fmt.Println(user.login)
// 	user, err = getUserByName(db, "Varabe")
// 	fmt.Println(user)
// }
