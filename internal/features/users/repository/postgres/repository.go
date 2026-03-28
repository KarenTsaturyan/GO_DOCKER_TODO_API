package users_postgres_repository

import core_postgres_pull "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/repository/postgres/pull"

type UsersRepository struct {
	pool core_postgres_pull.Pool
}

func NewUsersRepository(
	pool core_postgres_pull.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
