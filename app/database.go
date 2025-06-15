package app

import (
	"belajar-golang-retsful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:acumalaka@tcp(localhost:3306)/belajar_golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate -database "mysql://root:acumalaka@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:acumalaka@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down
	// migrate -database "mysql://root:acumalaka@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations version
	// migrate -database "mysql://root:acumalaka@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20250615082504
}
