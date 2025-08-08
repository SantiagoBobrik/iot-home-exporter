package main

import (
	"SantiagoBobrik/iot-home/config"
	"SantiagoBobrik/iot-home/db"
	"SantiagoBobrik/iot-home/domain"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	database := initDatabase()
	defer database.Close()

	app := fiber.New()

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	app.Post("/data", func(c fiber.Ctx) error {

		body := new(domain.Data)

		if err := c.Bind().JSON(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseError{
				Code:    400,
				Message: "invalid JSON body",
			})
		}

		if err := body.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseError{
				Code:    400,
				Message: err.Error(),
			})
		}

		config.GaugeMetrics.Temperature.WithLabelValues(body.DeviceID).Set(body.Temperature)
		config.GaugeMetrics.Humidity.WithLabelValues(body.DeviceID).Set(body.Humidity)

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

	log.Fatal(app.Listen(":3000"))

}

func initDatabase() *db.DB {
	database, err := db.New("./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	if err := database.InitSchema(); err != nil {
		log.Fatal(err)
	}

	return database
}

func init() {
	prometheus.MustRegister(config.GaugeMetrics.Humidity)
	prometheus.MustRegister(config.GaugeMetrics.Humidity)
}
