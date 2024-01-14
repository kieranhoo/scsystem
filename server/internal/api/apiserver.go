package api

import (
	"fmt"
	"log"
	"os"
	"scsystem/internal/config"

	"github.com/gofiber/fiber/v2"
)

type _App struct {
	engine *fiber.App
}

type IApp interface {
	Middleware(...func(*fiber.App)) IApp
	Route(...func(*fiber.App)) IApp
	Run() error
	Shutdown(<-chan os.Signal) IApp
	BackgroundTask(...func()) IApp
}

func NewServer() IApp {
	return &_App{
		engine: fiber.New(config.FiberConfig()),
	}
}

func (app *_App) BackgroundTask(tasks ...func()) IApp {
	for _, task := range tasks {
		go task()
	}
	return app
}

func (app *_App) Middleware(middlewares ...func(*fiber.App)) IApp {
	for _, middleware := range middlewares {
		middleware(app.engine)
	}
	return app
}

func (app *_App) Route(routes ...func(*fiber.App)) IApp {
	for _, route := range routes {
		route(app.engine)
	}
	return app
}

func (app *_App) Shutdown(sig <-chan os.Signal) IApp {
	go func() {
		<-sig
		fmt.Println()
		//cache.Shutdown()
		if config.StageStatus == "prod" {
			log.Println("[SERVER] Server is shutting down ..")
			if err := app.engine.Shutdown(); err != nil {
				log.Printf("Oops... Server is not shutting down! Reason: %v", err)
			}
		} else {
			os.Exit(0)
		}
	}()
	return app
}
