package users_postgres_repository

import "github.com/KarenTsaturyan/GO_DOCKER_TODO_API/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, u := range users {
		userDomains[i] = domain.NewUser(
			u.ID,
			u.Version,
			u.FullName,
			u.PhoneNumber,
		)
	}

	return userDomains
}
