package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thd3r/employee_management/internal/server"
	"github.com/thd3r/employee_management/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := server.New()
	routes.App(server)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
