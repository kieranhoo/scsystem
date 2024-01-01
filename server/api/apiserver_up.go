package api

import (
	"scsystem/api/middleware"
	"scsystem/api/routes"
	"scsystem/internal/config"
	"scsystem/internal/cron"
	"scsystem/internal/tasks"
	"scsystem/pkg/job"
	"scsystem/pkg/mailers"
	"scsystem/pkg/sentry"
	"scsystem/pkg/utils"
	"scsystem/pkg/worker"
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

func (app *_App) Scheduler() {
	newJob := job.New()
	go newJob.Scheduler(cron.ScanInValidDate, time.Hour*24*7)

	newJob.Info()
}

func (app *_App) Run() error {

	app.BackgroundTask()
	app.Scheduler()

	app.Middleware(
		middleware.FiberMiddleware,
		middleware.SentryMiddleware,
	)

	app.Route(
		routes.HealthCheck,
		routes.Auth,
		routes.Room,
		routes.User,
		routes.Stat,
		routes.Admin,
		routes.Swagger,
		routes.NotFoundRoute,
	)

	return utils.StartServer(app.engine)
}
