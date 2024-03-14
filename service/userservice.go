package service

import (
	"context"
	"learngo/helpers"
	"learngo/models"
	"learngo/settings"
	"learngo/store"
)

func UserListService(ctx context.Context) ([]models.User, error) {
	result, err := store.UserListStore(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UserService(ctx context.Context, id int64) (*models.User, error) {
	obj, err := store.UserGetByIDStore(ctx, id)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func UserCreateService(ctx context.Context, req models.User) (*models.User, error) {
	// connect to dbstore
	tx, err := settings.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer settings.RollbackTx(ctx, tx)

	//Create Empty variable
	req.PasswordHash = helpers.GetMd5(req.PasswordHash)

	obj, err := store.UserInsertStore(ctx, tx, req)
	if err != nil {
		return nil, err
	}

	if err := settings.CommitTx(ctx, tx); err != nil {
		return nil, err
	}

	return obj, nil
}
