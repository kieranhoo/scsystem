package api

import (
	"fmt"
	"log"
	"os"
	"qrcheckin/internal/config"
	"qrcheckin/pkg/utils"
	"qrcheckin/pkg/x/worker"

	"github.com/gofiber/fiber/v2"
)

type _App struct {
	engine *fiber.App
}

type IApp interface {
	Worker(*worker.Config) IApp
	Middleware(...middleware) IApp
	Route(...route) IApp
	Run() IApp
	Shutdown(<-chan os.Signal) IApp
	BackgroundTask(...backgroundTask) IApp
}

type middleware func(*fiber.App)
type route func(*fiber.App)
type backgroundTask func()

func New() IApp {
	return &_App{
		engine: fiber.New(config.FiberConfig()),
	}
}

func (app *_App) BackgroundTask(tasks ...backgroundTask) IApp {
	for _, task := range tasks {
		go task()
	}
	return app
}

// Worker
// Deprecated
func (app *_App) Worker(wcf *worker.Config) IApp {
	worker.Conf = wcf
	return app
}

func (app *_App) Middleware(middlewares ...middleware) IApp {
	for _, middleware := range middlewares {
		middleware(app.engine)
	}
	return app
}

func (app *_App) Route(routes ...route) IApp {
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

func (app *_App) Run() IApp {
	utils.StartServer(app.engine)
	return app
}
