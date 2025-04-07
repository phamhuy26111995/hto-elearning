package database

import "log"

func createUsersTable() {
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

func createCoursesTable() {
	query := `
		CREATE TABLE IF NOT EXISTS courses (
			 course_id SERIAL PRIMARY KEY,
			  title VARCHAR(255) NOT NULL,
			  description TEXT,
			  teacher_id INTEGER NOT NULL,  
			  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
