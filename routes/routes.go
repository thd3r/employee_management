package routes

import (
	"fmt"
	"os"

	"github.com/thd3r/employee_management/controllers"
	"github.com/thd3r/employee_management/internal/server"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var port = os.Getenv("PORT")

func App(s *server.FiberServer) {
	s.App.Get("/", controllers.Index)
	s.App.Get("/health", controllers.Health)

	api := s.App.Group("/api")
	api.Get("/", controllers.IndexApiHandler)

	v1 := api.Group("/v1")
	v1.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("http://localhost:%v", port),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	v1.Get("/", controllers.IndexApiHandler)
	v1.Get("/sitemap", controllers.ApiSitemap)

	employe := v1.Group("/employe")

	employe.Get("/", controllers.GetAllEmployeHandler)
	employe.Get("/:employeId/detail", controllers.GetEmployeByIdHandler)
	employe.Post("/create", controllers.CreateEmployeHandler)
	employe.Put("/:employeId/update", controllers.UpdateEmployeHandler)
	employe.Delete("/:employeId/delete", controllers.DeleteEmployeHandler)
}
