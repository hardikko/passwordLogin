package service

import (
	"context"
	"fmt"
	"learngo/helpers"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
)

func LoginCreateService(ctx context.Context, req models.Login) (*models.User, error) {
	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	//Check if user exists in db
	obj, err := store.GetUsersByEmailStore(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	req.PasswordHash = helpers.GetMd5(req.PasswordHash)

	// Compare User pass with req pass
	if req.PasswordHash != obj.PasswordHash {
		return nil, fmt.Errorf("invalid password")
	}

	if req.Email != obj.Email {
		return nil, fmt.Errorf("invalid Email")
	}
	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}
