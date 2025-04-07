package database

import (
	"database/sql"
	"fmt"
	"github.com/phamhuy26111995/hto-elearning/internal/config"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		return fmt.Errorf("failed to load database config: %w", err)
	}

	// Mở kết nối
	DB, err = sql.Open("postgres", dbConfig.ConnectionString())
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Cấu hình connection pool
	DB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	DB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	DB.SetConnMaxLifetime(time.Hour)
	DB.SetConnMaxIdleTime(time.Minute * 30)

	// Kiểm tra kết nối
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL with lib/pq!")
	return nil
}

//func CloseDB() {
//	if DB != nil {
//		DB.Close()
//		log.Println("Closed database connection")
//	}
//}

func InitDB() {
	err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	createUserTable()
}

func createUserTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			user_id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL, 
			role VARCHAR(20) DEFAULT 'STUDENT',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
