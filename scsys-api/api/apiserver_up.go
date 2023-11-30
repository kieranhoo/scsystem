package api

import (
	"scsystem/api/middleware"
	"scsystem/api/routes"
	"scsystem/internal/config"
	"scsystem/internal/service"
	"scsystem/internal/tasks"
	"scsystem/pkg/sentry"
	"scsystem/pkg/utils"
	"scsystem/pkg/x/job"
	"scsystem/pkg/x/mailers"
	"scsystem/pkg/x/worker"
	"time"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

func (w *Worker) Run() error {
	w.engine.HandleFunctions(tasks.Path())
	return w.engine.Run()
}

func JobLaunch() {
	j := job.New()
	j.Scheduler(service.Ping, 5*time.Second)

	j.Launch()
}

func (app *_App) Run() error {

	app.BackgroundTask(JobLaunch)

	app.Middleware(
		middleware.FiberMiddleware,
		middleware.SentryMiddleware,
	)

	app.Route(
		routes.Gateway,
		routes.NotFoundRoute,
	)

	return utils.StartServer(app.engine)
}
