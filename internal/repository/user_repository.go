package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/doffy007/user-login-register/config"
	"github.com/doffy007/user-login-register/database/query"
	"github.com/doffy007/user-login-register/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(entity.CreateUser) error
	FindOneUser(entity.FindOneUser, []string) (*entity.Users, error)
}

type userRepository struct {
	ctx    context.Context
	config *config.Config
	db     *sqlx.DB
}

func (r repository) UserRepository() UserRepository {
	return &userRepository{
		ctx:    r.ctx,
		config: r.config,
		db:     r.db,
	}
}

// CreateUser implements UserRepository.
func (r userRepository) CreateUser(payload entity.CreateUser) error {
	_, err := r.db.NamedExecContext(r.ctx, query.CreateUser, payload)

	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository CreateUser: %v\n", err.Error())
		return err
	}

	return nil
}

// FindOneUser implements UserRepository.
func (r userRepository) FindOneUser(payload entity.FindOneUser, fields []string) (*entity.Users, error) {
	var columns string
	if len(fields) == 0 {
		columns = "*"
	} else {
		columns = strings.Join(fields, ", ")
	}
	dest := entity.Users{}

	var filterValues []interface{}
	var filterKeys []string

	if payload.ID != nil {
		filterKeys = append(filterKeys, "id = ?")
		filterValues = append(filterValues, string(rune(*payload.ID)))
	}
	if payload.Email != nil {
		filterKeys = append(filterKeys, "email = ?")
		filterValues = append(filterValues, *payload.Email)
	}
	if payload.Username != nil {
		filterKeys = append(filterKeys, "username = ?")
		filterValues = append(filterValues, *payload.Username)
	}

	query := fmt.Sprintf(query.FindOneUser, columns, strings.Join(filterKeys, " AND "))
	err := r.db.GetContext(r.ctx, &dest, query, filterValues...)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Printf("\033[1;31m [ERROR] \033[0m Repository FindOneUser: %v\n", err.Error())
		return nil, err
	}

	return &dest, nil
}
