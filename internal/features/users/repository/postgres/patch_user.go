package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/domain"
	core_errors "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/errors"
	core_postgres_pool "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
        UPDATE todoapp.users
        SET
            full_name = $1,
            phone_number = $2,
            version = version + 1
        WHERE id = $3 AND version = $4
        RETURNING id, version, full_name, phone_number
    `

	var m UserModel
	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, id, user.Version)
	if err := row.Scan(
		&m.ID,
		&m.Version,
		&m.FullName,
		&m.PhoneNumber,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id '%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}

		return domain.User{}, fmt.Errorf("scan patched user: %w", err)
	}

	userDomain := domain.NewUser(
		m.ID,
		m.Version,
		m.FullName,
		m.PhoneNumber,
	)

	return userDomain, nil
}
