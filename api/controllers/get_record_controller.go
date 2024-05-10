package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golangcimri/api/globals"
	"github.com/golangcimri/api/models"
)

func GetRecord(c *fiber.Ctx) error {
	// Get the record id from the request
	id := c.Params("id")

	// ID'yi tamsayıya çevir
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Get the db connection
	db := globals.Variables.Database
	result := db.GetOne(db.Collection.Records, intID)
	if result.Err() != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Record not found",
		})
	}

	// Decode the record
	var record models.Record
	if err := result.Decode(&record); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error decoding record"})
	}

	// Return the record
	return c.JSON(record)

}
