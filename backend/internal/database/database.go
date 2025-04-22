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
	dbConfig := config.LoadSupabaseConfig()

	var err error
	// Mở kết nối
	DB, err = sql.Open("postgres", dbConfig)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Cấu hình connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
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

func InitDB() {
	err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	createUsersTable()
	createCoursesTable()
	createModulesTable()
	createLessonsTable()
}
