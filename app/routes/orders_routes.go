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
	app.Get("/item/:id", func(c *fiber.Ctx) error {
		items := models.Items

		itemIdParam := c.Params("id")
		itemId, err := strconv.Atoi(itemIdParam)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  err.Error(),
			})
		}
		var itemToSend models.Item

		for index, item := range items {
			if item.ID == itemId {
				itemToSend = models.Items[index]
			}
		}

		return c.JSON(fiber.Map{
			"status": c.Response().StatusCode(),
			"data":   itemToSend,
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
