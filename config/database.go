package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Ganti driver SQLite
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "./pharmacy.db")
	if err != nil {
		log.Fatal("Error saat membuka koneksi database:", err)
	}

	query := `
    CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT,
        price REAL NOT NULL,
        stock INTEGER NOT NULL
    );
    `
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Error saat membuat tabel:", err)
	}

	log.Println("Database dan tabel berhasil diinisialisasi.")
}
