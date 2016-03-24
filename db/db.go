package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath string = "./hits.db"

type Log struct {
	Id        int
	Path      string
	UserAgent string
}
type Logs []Log

type Path string
type Paths []Path

func InitializeDb() {

	os.Remove(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
  create table hits (id integer not null primary key, path text, user_agent text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func InsertPath(path string, userAgent string) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into hits(path, user_agent) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(path, userAgent)
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
			user_agent string
		)
		err := rows.Scan(&id, &path, &user_agent)
		if err != nil {
			log.Fatal(err)
		}

		logs = append(logs, Log{id, path, user_agent})
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
			path Path
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
