package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func RegisterRoutes() *gin.Engine {

	server := gin.Default()

	userController := registerUserController()
	courseController := registerCourseController()
	moduleController := registerModuleController()
	lessonController := registerLessonController()
	quizController := registerQuizController()
	quizQuestionController := registerQuizQuestionController()
	quizOptionController := registerQuizOptionController()

	authenticated := server.Group("/api/v1/teacher")

	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/users", userController.GetUsers)
	authenticated.POST("/users/create", userController.CreateUser)
	authenticated.POST("/users/create-student", userController.CreateStudent)
	authenticated.PUT("/users/update", userController.UpdateUser)
	authenticated.GET("/users/:id", userController.GetUserById)

	authenticated.GET("/courses", courseController.GetAllCourses)
	authenticated.POST("/course/create", courseController.CreateCourse)
	authenticated.PUT("/course/update", courseController.UpdateCourse)
	authenticated.GET("/course/:id", courseController.GetCourse)
	authenticated.DELETE("/course/:id", courseController.DeleteCourse)

	authenticated.GET("/modules", moduleController.GetAllModulesByCourse)
	authenticated.POST("/modules/create", moduleController.CreateModules)
	authenticated.PUT("/modules/update", moduleController.UpdateModules)

	authenticated.GET("/lessons", lessonController.GetLessonsByModuleId)
	authenticated.POST("/lessons/create", lessonController.CreateLessons)
	authenticated.PUT("/lessons/update", lessonController.UpdateLessons)

	authenticated.GET("/quizzes", quizController.GetQuizzesByModuleId)
	authenticated.POST("/quizzes/create", quizController.CreateQuizzes)
	authenticated.PUT("/quizzes/update", quizController.UpdateQuizzes)

	authenticated.GET("/quiz-questions", quizQuestionController.GetAllQuestionsByQuizId)
	authenticated.POST("/quiz-questions/create", quizQuestionController.CreateQuestions)
	authenticated.PUT("/quiz-questions/update", quizQuestionController.UpdateQuestions)

	authenticated.GET("/quiz-options", quizOptionController.GetAllQuizOptionsByQuestionId)
	authenticated.POST("/quiz-options/create", quizOptionController.CreateQuizOptions)
	authenticated.PUT("/quiz-options/update", quizOptionController.UpdateQuizOptions)

	server.POST("/login", userController.Login)

	return server
}
