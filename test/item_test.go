package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetItemById(t *testing.T) {
	app := fiber.New()

	request := httptest.NewRequest("POST", "localhost:8900/item/1023", nil)
	response, err := app.Test(request)

	fmt.Println(response)

	assert.Nil(t, err)
}
