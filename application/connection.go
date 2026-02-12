package application

import (
	"database/sql"
	"fmt"
	"log"
	"module/portofolio1/database/migrations"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type DatabaseConfig struct {
	*sql.DB
}

func NewConnection() error{
	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/mydatabase"

	var err error

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("gagal membuka koneksi database: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetConnMaxIdleTime(5 * time.Minute)

	// Test koneksi
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("gagal ping database: %w", err)
	}

	log.Println("Koneksi ke MySQL berhasil!")
	return nil
}

func CreateTable(db *DatabaseConfig)error {
	for _, query := range migrations.AllMigrations {
		if _,err := db.Exec(query); err != nil {
			fmt.Println("Gagal membuat tabel:", err)
			return err
		}
	}
	return nil
}