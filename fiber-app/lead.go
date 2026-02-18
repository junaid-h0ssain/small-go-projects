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

func getLead(c *fiber.Ctx) error {

}

func getLeads(c *fiber.Ctx) error {

}

func newLead(c *fiber.Ctx) error {

}

func deleteLead(c *fiber.Ctx) error {

}


