package database

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "./trustcoupon.db")
    if err != nil {
        log.Fatal("Failed to open database:", err)
    }

    // 建立資料表
    createTables()
    log.Println("✅ Database initialized")
}

func CloseDB() {
    if DB != nil {
        DB.Close()
    }
}

func createTables() {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS transactions (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            transaction_id TEXT UNIQUE NOT NULL,
            customer_name TEXT NOT NULL,
            discount_amount INTEGER NOT NULL,
            expired_date TEXT NOT NULL,
            qr_code TEXT,
            deep_link TEXT,
            status TEXT DEFAULT 'pending',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
        `CREATE TABLE IF NOT EXISTS coupons (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            transaction_id TEXT UNIQUE NOT NULL,
            cid TEXT UNIQUE,
            customer_name TEXT NOT NULL,
            discount_amount INTEGER NOT NULL,
            expired_date TEXT NOT NULL,
            is_used BOOLEAN DEFAULT FALSE,
            used_at DATETIME,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id)
        )`,
        `CREATE TABLE IF NOT EXISTS verifications (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            verification_id TEXT UNIQUE NOT NULL,
            customer_name TEXT,
            discount_amount INTEGER,
            expired_date TEXT,
            original_amount DECIMAL(10,2),
            final_amount DECIMAL(10,2),
            verified_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
    }

    for _, query := range queries {
        _, err := DB.Exec(query)
        if err != nil {
            log.Fatal("Failed to create table:", err)
        }
    }
}