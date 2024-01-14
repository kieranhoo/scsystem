package router

import (
	"scsystem/internal/messaging/handler"
	"scsystem/internal/messaging/queue"
	"scsystem/pkg/worker"
)

func Common(eng *worker.Engine) {
	eng.Queue(queue.WorkerHealthCheck, handler.HandleHealthCheck)
}
