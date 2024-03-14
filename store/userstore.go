package store

import (
	"context"
	"learngo/models"
	"learngo/settings"

	"github.com/jackc/pgx/v4"
)

func UserListStore(ctx context.Context) ([]models.User, error) {
	result := []models.User{}
	obj := models.User{}

	queryStmt := `SELECT * FROM users`
	rows, err := settings.DBClient.Query(ctx, queryStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(
			&obj.ID,
			&obj.FirstName,
			&obj.LastName,
			&obj.Email,
			&obj.PasswordHash,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, obj)
	}

	return result, nil
}
func UserGetByIDStore(ctx context.Context, id int64) (*models.User, error) {
	obj := models.User{}

	queryStmt := `
		SELECT * FROM users
		WHERE users.id = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, id)
	if err := row.Scan(
		&obj.ID,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.PasswordHash,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &obj, nil
}

func UserInsertStore(ctx context.Context, tx pgx.Tx, arg models.User) (*models.User, error) {
	obj := &models.User{}

	queryStmt := `
	INSERT INTO
	users(
		first_name,
		last_name,
		email,
		password_hash

	)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	row := tx.QueryRow(ctx, queryStmt,
		&arg.FirstName,
		&arg.LastName,
		&arg.Email,
		&arg.PasswordHash,
	)

	if err := row.Scan(
		&obj.ID,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.PasswordHash,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return obj, nil
}

func GetUsersByEmailStore(ctx context.Context, email string) (*models.User, error) {
	obj := models.User{}

	queryStmt := `
		SELECT * FROM users
		WHERE users.email = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, email)
	if err := row.Scan(
		&obj.ID,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.PasswordHash,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &obj, nil
}

func GetUsersByPasswordStore(ctx context.Context, password string) (*models.User, error) {
	obj := models.User{}

	queryStmt := `
		SELECT * FROM users
		WHERE users.password_hash = $1
	`
	row := settings.DBClient.QueryRow(ctx, queryStmt, password)
	if err := row.Scan(
		&obj.ID,
		&obj.FirstName,
		&obj.LastName,
		&obj.Email,
		&obj.PasswordHash,
		&obj.CreatedAt,
		&obj.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &obj, nil
}
