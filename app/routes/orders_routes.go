package routes

import (
	"appbrew/app/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func OrdersRoutes(app *fiber.App) {
	app.Get("/orders", func(c *fiber.Ctx) error {
		items := models.Items

		return c.Render("orders_page", fiber.Map{
			"items": items,
		})
	})
	app.Post("/orders/pay", func(c *fiber.Ctx) error {
		var formData map[string]int

		err := c.BodyParser(&formData)
		if err != nil {
			log.Println(err)
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		var total int64
		for itemID, quantity := range formData {
			for _, item := range models.Items {
				if strconv.Itoa(item.ID) == itemID {
					total += int64(quantity) * item.Price
				}
			}
		}

		fmt.Println(total)
		return c.JSON(fiber.Map{
			"status":  c.Response().StatusCode(),
			"message": "Pembayaran Sukses",
			"items":   formData,
			"total":   total,
		})
	})

}
