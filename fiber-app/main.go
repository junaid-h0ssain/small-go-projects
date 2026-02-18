package main

import (
	"fmt"

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
		fmt.Printf("Database connection error: %v\n", err)
		panic("failed to connect database")
	}
	defer func() {
		_ = db.Close()
	}()

	db.AutoMigrate(&Lead{})
	DBConn = db

	app := fiber.New()
	setupRoutes(app)
	fmt.Println("Listening on port 8080")
	_ = app.Listen(8080)

}
