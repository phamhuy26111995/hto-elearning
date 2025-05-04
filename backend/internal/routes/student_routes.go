package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func StudentRoutes(server *gin.Engine) {

	userController := RegisterUserController()
	courseController := RegisterCourseController()
	moduleController := RegisterModuleController()
	lessonController := RegisterLessonController()
	quizController := RegisterQuizController()
	quizQuestionController := RegisterQuizQuestionController()
	quizOptionController := RegisterQuizOptionController()

	authenticated := server.Group("/api/v1/student")

	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/users/:id", userController.GetUserById)

	authenticated.GET("/courses", courseController.GetAllCourses)
	authenticated.GET("/course/:id", courseController.GetCourse)

	authenticated.GET("/modules", moduleController.GetAllModulesByCourse)

	authenticated.GET("/lessons/:moduleId", lessonController.GetLessonsByModuleId)

	authenticated.GET("/quizzes", quizController.GetQuizzesByModuleId)

	authenticated.GET("/quiz-questions", quizQuestionController.GetAllQuestionsByQuizId)

	authenticated.GET("/quiz-options", quizOptionController.GetAllQuizOptionsByQuestionId)
}
