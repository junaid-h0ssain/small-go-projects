package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", getLeads)
	app.Get("/:id", getLead)
	app.Post("/", newLead)
	app.Delete("/", deleteLead)
}

func main() {
	app := fiber.New()

}
