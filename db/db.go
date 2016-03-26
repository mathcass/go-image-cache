package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath string = "./hits.db"

type Log struct {
	Id         int
	Path       string
	InsertTime string
	UserAgent  string
}
type Logs []Log

// For some reasoning, type aliasing string as Path breaks the database Scan
type Paths []string

func InitializeDb() {
	if _, err := os.Stat(dbPath); err == nil {
		// File exists but verify that we can open it properly
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		return

	} else {
		// File does not exist, create it and initialize db
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		sqlStmt := `
  create table hits (id integer not null primary key, path text, time text, user_agent text);
	`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return
		}
	}

}

func InsertPath(path string, userAgent string) {
	insertTime := time.Now().String() // Insert the current time

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into hits(path, time, user_agent) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(path, insertTime, userAgent)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}

func GetPathResults(path string) Logs {
	var logs Logs

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from hits where path = ?", path)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         int
			path       string
			insertTime string
			userAgent  string
		)
		err := rows.Scan(&id, &path, &insertTime, &userAgent)
		if err != nil {
			log.Fatal(err)
		}

		logs = append(logs, Log{id, path, insertTime, userAgent})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return logs
}

func GetUniquePathResults() Paths {
	var paths Paths

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select distinct path from hits")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			path string
		)
		err := rows.Scan(&path)
		if err != nil {
			log.Fatal(err)
		}
		paths = append(paths, path)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return paths
}
