package main

import (
	"log"

	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/db"
	"go-fiber-api-docker/pkg/products"
	"go-fiber-api-docker/pkg/suppliers"
	"go-fiber-api-docker/pkg/users"
	"go-fiber-api-docker/pkg/common/auth"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)
	suppliers.RegisterRoutes(app, h)
	users.RegisterRoutes(app, h)
	auth.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
