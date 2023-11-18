package worker

import (
	"github.com/RichardKnop/machinery/v1/config"
)

// Worker
// Deprecated
type Worker struct {
	Task   map[string]interface{}
	Config *config.Config
}

// Tasks
// Deprecated
type Tasks map[string]Task

// Task
// Deprecated
type Task map[string]interface{}

// Config
// Deprecated
type Config struct {
	Broker        string
	ResultBackend string
	Tasks         Tasks
}

// NewWorker
// Deprecated: use NewServer() instead
func NewWorker(queueName string, wcf *Config) *Worker {
	return &Worker{
		Config: &config.Config{
			Broker:          wcf.Broker,
			DefaultQueue:    queueName,
			ResultBackend:   wcf.ResultBackend,
			ResultsExpireIn: 3600,
			// AMQP: &config.AMQPConfig{
			// 	Exchange:      "qrcheckin_exchange",
			// 	ExchangeType:  "direct",
			// 	BindingKey:    queueName,
			// 	PrefetchCount: 3,
			// },
			Redis: &config.RedisConfig{
				MaxIdle:                3,
				IdleTimeout:            240,
				ReadTimeout:            15,
				WriteTimeout:           15,
				ConnectTimeout:         15,
				NormalTasksPollPeriod:  1000,
				DelayedTasksPollPeriod: 500,
			},
		},
		Task: wcf.Tasks[queueName],
	}
}
