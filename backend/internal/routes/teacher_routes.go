package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phamhuy26111995/hto-elearning/internal/middlewares"
)

func TeacherRoutes(server *gin.Engine) {

	userController := RegisterUserController()
	courseController := RegisterCourseController()
	moduleController := RegisterModuleController()
	lessonController := RegisterLessonController()
	quizController := RegisterQuizController()
	quizQuestionController := RegisterQuizQuestionController()
	quizOptionController := RegisterQuizOptionController()

	authenticated := server.Group("/api/v1/teacher")

	authenticated.Use(middlewares.Authenticate, middlewares.AuthorizeTeacher)

	authenticated.GET("/users", userController.GetUsers)
	authenticated.GET("/users/get-students", userController.GetUsersByTeacher)
	authenticated.POST("/users/create-student", userController.CreateStudent)
	authenticated.PUT("/users/update", userController.UpdateUser)
	authenticated.PUT("/users/update-student", userController.UpdateUser)
	authenticated.POST("/users/enrollment", userController.EnrollCourseForStudent)
	authenticated.DELETE("/users/unenrollment", userController.UnEnrollCourseForStudent)
	authenticated.POST("/users/change-status", userController.ChangeStatus)
	authenticated.GET("/users/current-user", userController.GetCurrentUserLogin)
	authenticated.GET("/users/:id", userController.GetUserById)

	authenticated.GET("/courses", courseController.GetAllCourses)
	authenticated.POST("/course/create", courseController.CreateCourse)
	authenticated.PUT("/course/update", courseController.UpdateCourse)
	authenticated.GET("/course/:id", courseController.GetCourse)
	authenticated.DELETE("/course/:id", courseController.DeleteCourse)

	authenticated.GET("/modules/:courseId", moduleController.GetAllModulesByCourse)
	authenticated.POST("/modules/create", moduleController.CreateModules)
	authenticated.PUT("/modules/update", moduleController.UpdateModules)

	authenticated.GET("/lessons/:moduleId", lessonController.GetLessonsByModuleId)
	authenticated.POST("/lessons/create", lessonController.CreateLessons)
	authenticated.PUT("/lessons/update", lessonController.UpdateLessons)

	authenticated.GET("/quizzes/:moduleId", quizController.GetQuizzesByModuleId)
	authenticated.POST("/quizzes/create", quizController.CreateQuizzes)
	authenticated.PUT("/quizzes/update", quizController.UpdateQuizzes)

	authenticated.GET("/quiz-questions/:quizId", quizQuestionController.GetAllQuestionsByQuizId)
	authenticated.POST("/quiz-questions/create", quizQuestionController.CreateQuestions)
	authenticated.PUT("/quiz-questions/update", quizQuestionController.UpdateQuestions)

	authenticated.GET("/quiz-options/:questionId", quizOptionController.GetAllQuizOptionsByQuestionId)
	authenticated.POST("/quiz-options/create", quizOptionController.CreateQuizOptions)
	authenticated.PUT("/quiz-options/update", quizOptionController.UpdateQuizOptions)
}
