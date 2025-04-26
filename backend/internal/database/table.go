package database

import "log"

func createUsersTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			user_id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
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

func createModulesTable() {
	query := `
		CREATE TABLE IF NOT EXISTS modules (
				module_id SERIAL PRIMARY KEY,
			  course_id INTEGER NOT NULL,    
			  title VARCHAR(255) NOT NULL,
			  description TEXT,
			  order_index INTEGER DEFAULT 1
				);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createLessonsTable() {
	query := `
		CREATE TABLE IF NOT EXISTS lessons  (
		  lesson_id SERIAL PRIMARY KEY,
		  module_id INTEGER NOT NULL,             
		  title VARCHAR(255) NOT NULL,
		  content TEXT,
		  order_index INTEGER DEFAULT 1,
		  video_url VARCHAR(255),
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

func createQuizzesTable() {
	query := `
		CREATE TABLE IF NOT EXISTS quizzes   (
  			quiz_id SERIAL PRIMARY KEY,
  			module_id INTEGER, 
  			course_id INTEGER NOT NULL,    -- tham chiếu đến courses.course_id
 		 title VARCHAR(255) NOT NULL,
  		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createQuizQuestionsTable() {
	query := `
		CREATE TABLE IF NOT EXISTS quiz_questions   (
  			question_id  SERIAL PRIMARY KEY,
  			quiz_id INTEGER NOT NULL, 
  			question_content TEXT NOT NULL,    -- tham chiếu đến courses.course_id
 		 question_type VARCHAR(50) DEFAULT 'ONE_CHOICE'
				);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createQuizOptionsTable() {
	query := `
		CREATE TABLE IF NOT EXISTS quiz_options   (
  			option_id  SERIAL PRIMARY KEY,
  			question_id INTEGER NOT NULL, 
  			option_text TEXT NOT NULL,    -- tham chiếu đến courses.course_id
 		 is_correct BOOLEAN DEFAULT false
				);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
