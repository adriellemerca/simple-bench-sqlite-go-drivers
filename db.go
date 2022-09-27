package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	dbPath = "./test.db"
)

type database struct {
	sqlDB *sql.DB
}

func prepareDB(driver string) (db database) {
	newDb, err := sql.Open(driver, dbPath)
	if err != nil {
		log.Fatal(fmt.Errorf("open db error %v", err))
		return
	}

	db.sqlDB = newDb
	return
}

func (db database) closeDB() {
	if err := os.Remove(dbPath); err != nil {
		log.Fatal(fmt.Errorf("remove db file error: %v", err))
	}

	if err := db.sqlDB.Close(); err != nil {
		log.Fatal(fmt.Errorf("close db error: %v", err))
	}
}

func (db database) initialSeed() {
	createTableStmt, err := db.sqlDB.Prepare("CREATE TABLE tests (id INTEGER PRIMARY KEY AUTOINCREMENT, field TEXT)")
	if err != nil {
		log.Fatal(fmt.Errorf("prepare 'create table' db error: %v", err))
	}
	defer func() {
		if err := createTableStmt.Close(); err != nil {
			log.Fatal(fmt.Errorf("close 'create table' db stmt error: %v", err))
		}
	}()

	if _, err := createTableStmt.Exec(); err != nil {
		log.Fatal(fmt.Errorf("exec 'create table' db stmt error: %v", err))
	}
}

func (db database) prepareInsertStmt() *sql.Stmt {
	insertRowStmt, err := db.sqlDB.Prepare("INSERT INTO tests (field) VALUES ('test')")
	if err != nil {
		log.Fatal(fmt.Errorf("prepare 'create table' db error: %v", err))
	}
	return insertRowStmt
}
