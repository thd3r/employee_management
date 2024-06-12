package routes

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thd3r/employee_management/controllers"
	"github.com/thd3r/employee_management/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/joho/godotenv/autoload"
)

var port, _ = strconv.Atoi(os.Getenv("PORT"))

func App(s *server.FiberServer) {
	s.App.Get("/", controllers.Index)
	s.App.Get("/health", controllers.Health)

	api := s.App.Group("/api")
	api.Get("/", controllers.IndexApiHandler)

	v1 := api.Group("/v1")
	v1.Get("/", controllers.IndexApiHandler)
	v1.Get("/sitemap", controllers.ApiSitemap)

	employee := v1.Group("/employee")
	employee.Use(func(c *fiber.Ctx) error {
		c.Request().Header.Set("Origin", fmt.Sprintf("http://localhost:%d", port))
		return c.Next()
	})

	employee.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("http://localhost:%d", port),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	employee.Get("/", controllers.GetAllEmployeeHandler)
	employee.Get("/:employeeId/detail", controllers.GetEmployeeByIdHandler)
	employee.Post("/create", controllers.CreateEmployeeHandler)
	employee.Put("/:employeeId/update", controllers.UpdateEmployeeHandler)
	employee.Delete("/:employeeId/delete", controllers.DeleteEmployeeHandler)
}
