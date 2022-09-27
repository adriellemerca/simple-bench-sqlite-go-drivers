package main

import (
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func BenchmarkInsertRowsCGO(b *testing.B) {
	db := prepareDB("sqlite3")
	defer db.closeDB()
	db.initialSeed()

	insertRowStmt := db.prepareInsertStmt()
	defer func() {
		if err := insertRowStmt.Close(); err != nil {
			log.Fatal(fmt.Errorf("close 'insert row' db stmt error: %v", err))
		}
	}()

	for i := 0; i < b.N; i++ {
		if _, err := insertRowStmt.Exec(); err != nil {
			fmt.Printf("exec 'insert row' db stmt error: %v", err)
			b.Fail()
		}
	}
}

func BenchmarkSelectRowsCGO(b *testing.B) {
	db := prepareDB("sqlite3")
	defer db.closeDB()
	db.initialSeed()

	insertRowStmt := db.prepareInsertStmt()
	defer func() {
		if err := insertRowStmt.Close(); err != nil {
			log.Fatal(fmt.Errorf("close 'insert row' db stmt error: %v", err))
		}
	}()

	for i := 0; i < 100; i++ {
		if _, err := insertRowStmt.Exec(); err != nil {
			fmt.Printf("exec 'insert row' db stmt error: %v", err)
			b.Fail()
		}
	}

	for i := 0; i < b.N; i++ {
		if _, err := db.sqlDB.Query("select * from tests"); err != nil {
			fmt.Printf("exec 'query' db error: %v", err)
			b.Fail()
		}
	}
}
