package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get(
		"/", func(c fiber.Ctx) error {
			return c.SendString("Hello, World!")
		},
	)

	// TLS configuration
	listenConfig := fiber.ListenConfig{
		CertFile:    "cert.pem",
		CertKeyFile: "key.pem",
	}

	// Listen on port 443 with TLS
	log.Fatal(app.Listen(":443", listenConfig))
}
