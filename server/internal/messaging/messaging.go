package messaging

import (
	"scsystem/internal/messaging/queue"
	"scsystem/pkg/worker"
)

type Messaging struct {
	engine *worker.Engine
}

func NewMessaging() *Messaging {
	return &Messaging{
		engine: worker.NewServer(worker.Priority{
			queue.CriticalQueue: 6, // processed 60% of the time
			queue.DefaultQueue:  3, // processed 30% of the time
			queue.LowQueue:      1, // processed 10% of the time
		}),
	}
}

func (msg *Messaging) Queue(queue ...func(eng *worker.Engine)) {
	for _, q := range queue {
		q(msg.engine)
	}
}

func (msg *Messaging) Run(concurrency int) error {
	msg.Queue(queue.Common)
	
	return msg.engine.Run(concurrency)
}
