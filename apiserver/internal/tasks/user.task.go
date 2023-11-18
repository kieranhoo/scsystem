package tasks

import (
	"context"
	"encoding/json"
	"scsystem/internal/model"
	"scsystem/internal/repo"

	"github.com/hibiken/asynq"
)

func HandleSaveUser(_ context.Context, task *asynq.Task) error {
	var user model.Users
	if err := json.Unmarshal(task.Payload(), &user); err != nil {
		return err
	}
	return repo.NewUser().Insert(&user)
}
