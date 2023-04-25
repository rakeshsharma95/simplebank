package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/rdevelop/simplebank/db/sqlc"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskPorocessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(rediOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		rediOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
		},
	)

	return &RedisTaskPorocessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskPorocessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	return processor.server.Start(mux)
}
