package router

import (
	"scsystem/internal/broker/handler"
	"scsystem/internal/broker/queue"
	"scsystem/pkg/worker"
)

func Common(eng *worker.Engine) {
	eng.Queue(queue.WorkerHealthCheck, handler.HandleHealthCheck)
}
