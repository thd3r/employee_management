package routes

import (
	"github.com/thd3r/employee_management/controllers"
	"github.com/thd3r/employee_management/internal/server"
)

func App(s *server.FiberServer) {
	s.App.Get("/", controllers.Index)
	s.App.Get("/health", controllers.Health)

	v1 := s.App.Group("/api").Group("/v1")

	employe := v1.Group("/employe")
	employe.Get("/", controllers.GetAllEmployeHandler)
	employe.Get("/:employeId/detail", controllers.GetEmployeByIdHandler)
	employe.Post("/create", controllers.CreateEmployeHandler)
	employe.Put("/:employeId/update", controllers.UpdateEmployeHandler)
	employe.Delete("/:employeId/delete", controllers.DeleteEmployeHandler)
}
