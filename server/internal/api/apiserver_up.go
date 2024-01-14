package api

import (
	"scsystem/internal/api/middleware"
	"scsystem/internal/api/routes"
	"scsystem/internal/config"
	"scsystem/internal/misc/cron"
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
