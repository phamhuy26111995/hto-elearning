package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Name         string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
}

//func LoadDBConfig() (DBConfig, error) {
//	// Load .env file
//	err := godotenv.Load()
//	if err != nil {
//		return DBConfig{}, fmt.Errorf("error loading .env file: %w", err)
//	}
//
//	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
//	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
//
//	return DBConfig{
//		Host:         os.Getenv("DB_HOST"),
//		Port:         os.Getenv("DB_PORT"),
//		User:         os.Getenv("DB_USER"),
//		Password:     os.Getenv("DB_PASSWORD"),
//		Name:         os.Getenv("DB_NAME"),
//		SSLMode:      os.Getenv("DB_SSLMODE"),
//		MaxOpenConns: maxOpen,
//		MaxIdleConns: maxIdle,
//	}, nil
//}
//
//func (c DBConfig) ConnectionString() string {
//	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
//		c.Host,
//		c.Port,
//		c.User,
//		c.Password,
//		c.Name,
//		c.SSLMode,
//	)
//}

func LoadSupabaseConfig() string {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}

	return os.Getenv("SUPABASE_POSTGRESQL")
}
