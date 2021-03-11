package client

import (
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"context"
)

type AuthCheckerClient interface {
	Create(ctx context.Context, id string) (models.Result, error)
	Check(ctx context.Context, id string) (models.Result, error)
	Delete(ctx context.Context, id string) (models.Result, error)
}
