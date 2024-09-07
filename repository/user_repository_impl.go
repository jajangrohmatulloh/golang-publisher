package repository

import (
	"context"
	"database/sql"
	"errors"
	"publisher/model/domain"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	SQL := "select id, first_name, last_name, password from category where username = ?"
	rows, err := repository.DB.QueryContext(ctx, SQL, username)
	user := domain.User{}
	if err != nil {
		return user, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Password)
		return user, nil
	} else {
		return user, errors.New("Username " + username + " Not Found")
	}
}
