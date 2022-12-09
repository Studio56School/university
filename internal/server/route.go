package server

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) routeApiV1(r *echo.Echo) {
	apiv1 := r.Group("api/v1")
	apiv1.GET("/students", s.handlers.university.GetStudents)
	apiv1.GET("/students/:id", s.handlers.university.GetStudentsById)
	apiv1.POST("/students/create", s.handlers.university.CreateStudent)
	apiv1.DELETE("/students/:id", s.handlers.university.DeleteStudent)
}
