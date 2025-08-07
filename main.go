package main

import (
	"SantiagoBobrik/iot-home/db"
	"SantiagoBobrik/iot-home/domain"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database, err := db.New("./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := database.InitSchema(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/data", func(c fiber.Ctx) error {

		body := new(domain.Data)

		if err := c.Bind().JSON(body); err != nil {
			return err
		}

		body.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

		if err := database.InsertData(*body); err != nil {
			errResponse := domain.ResponseError{Code: 500, Message: "error on creating row data"}
			fmt.Printf("Error %v - %v", errResponse.Error(), err)
			return c.Status(500).JSON(errResponse)
		}
		return c.Status(fiber.StatusCreated).Send([]byte{})
	})

	app.Get("/data", func(c fiber.Ctx) error {
		data, err := database.GetData()
		if err != nil {
			errResponse := domain.ResponseError{Code: 500, Message: "error on getting data"}
			fmt.Printf("Error %v - %v", errResponse.Error(), err)
			return c.Status(500).JSON(errResponse)
		}
		return c.Status(fiber.StatusOK).JSON(domain.Response{Data: data})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
