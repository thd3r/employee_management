package tests

import (
	"io"
	"testing"
	"net/http"

	"github.com/thd3r/employee_management/controllers"
	"github.com/thd3r/employee_management/internal/server"
	"github.com/thd3r/employee_management/routes"

	"github.com/gofiber/fiber/v2"
)

func TestHandler(t *testing.T) {
	// Create a Fiber app for testing
	app := fiber.New()
	// Inject the Fiber app into the server
	s := &server.FiberServer{App: app}
	routes.App(s)
	// Define a route in the Fiber app
	app.Get("/", controllers.Index)
	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("error creating request. Err: %v", err)
	}
	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	// Your test assertions...
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"code\":200}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}
