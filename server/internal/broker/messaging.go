package broker

import (
	"scsystem/internal/broker/queue"
	"scsystem/internal/broker/router"
	"scsystem/pkg/worker"
)

type broker struct {
	engine *worker.Engine
}

func Newbroker() *broker {
	return &broker{
		engine: worker.NewServer(worker.Priority{
			queue.CriticalQueue: 6, // processed 60% of the time
			queue.DefaultQueue:  3, // processed 30% of the time
			queue.LowQueue:      1, // processed 10% of the time
		}),
	}
}

func (msg *broker) router(queue ...func(eng *worker.Engine)) {
	for _, q := range queue {
		q(msg.engine)
	}
}

func (msg *broker) Run(concurrency int) error {
	msg.router(router.Common)

	return msg.engine.Run(concurrency)
}
