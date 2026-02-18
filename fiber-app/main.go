package main

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		getLeads(c)
	})
	app.Get("/:id", func(c *fiber.Ctx) {
		getLead(c)
	})
	app.Post("/", func(c *fiber.Ctx) {
		newLead(c)
	})
	app.Delete("/", func(c *fiber.Ctx) {
		deleteLead(c)
	})
}

func main() {
	db, err := gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer func() {
		_ = db.Close()
	}()

	db.AutoMigrate(&Lead{})
	DBConn = db

	app := fiber.New()
	setupRoutes(app)
	_ = app.Listen(3000)
}
