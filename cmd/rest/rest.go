package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"stori/internal/config"
	database "stori/internal/database/application"
	"stori/pkg/processor/application"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

//go:embed all:assets/*
var assets embed.FS

func main() {
	app := fiber.New()

	db, err := database.New().Connect()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/:email", func(c *fiber.Ctx) error {
		_ = config.New(config.Cfg{FS: assets, FileName: "data/ex1.csv", Email: c.Params("email"), Database: db})

		res, err := application.New().Process()

		if err != nil {
			return c.SendString(fmt.Sprintf("error %v", err))
		}

		return c.SendString(fmt.Sprintf("%v", res))
	})

	go func() {
		app.Listen(":3000")
	}()

	r := recover()
	if r != nil {
		log.Fatal(r)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	<-stop

	log.Println("Stopp....")
	// Handle here
	_ = database.New().Disconnect()
}
