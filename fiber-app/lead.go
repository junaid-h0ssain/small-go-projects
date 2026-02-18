package main

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

func getLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := DBConn
	var lead Lead
	db.Find(&lead, id)
	_ = c.JSON(lead)
}

func getLeads(c *fiber.Ctx) {
	db := DBConn
	var leads []Lead
	db.Find(&leads)
	_ = c.JSON(leads)
}

func newLead(c *fiber.Ctx) {
	db := DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	_ = c.JSON(lead)
}

func deleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully Deleted")
}


