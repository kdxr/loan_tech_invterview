package main

import (
	"fmt"
	"log"
	"manage_payments/configs"
	"manage_payments/database"
	"manage_payments/middlewares"
	"manage_payments/routes"
	"manage_payments/services"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	*fiber.App
}

func init() {
	// Load .env
	configs.LoadConfig()
}

func main() {

	conf := configs.GetConfig()

	services.InitSocket()

	if err := database.New(conf.Database); err != nil {
		log.Fatalln(err)
	}

	app := App{
		fiber.New(fiber.Config{
			Prefork: conf.Service.Env == "production",
			// EnablePrintRoutes: true,
			Immutable: true,
			AppName:   "Loan Management",
		}),
	}

	app.Use(cors.New(), recover.New(), logger.New(), middlewares.LimitterMiddleWare())

	routes.InitRoutes(app.App)

	port := configs.Config.Service.ServerPort

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		app.exit()
	}()

	if configs.Config.Service.SSL == "TRUE" {
		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		fmt.Println("Server listening on port with ssl:", port)
		log.Fatal(app.ListenTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY))
	} else {
		fmt.Println("Server listening on port:", port)
		log.Fatal(app.Listen(":" + port))
	}

}

func (app *App) exit() {
	_ = app.Shutdown()
}
