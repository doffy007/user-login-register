package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doffy007/user-login-register/config"
	"github.com/doffy007/user-login-register/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := &userRepository{
		ctx:    context.TODO(),
		config: &config.ConfigureApp,
		db:     sqlx.NewDb(db, "sqlmock"),
	}

	mock.ExpectExec("INSERT INTO users").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateUser(entity.CreateUser{
		Username: new(string),
		Email:    new(string),
		Password: new(string),
	})

	assert.NoError(t, err, "Unexpected error")
	assert.NoError(t, mock.ExpectationsWereMet(), "Expectations not met")
}

func TestUserRepository_FindOneUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := &userRepository{
		ctx:    context.TODO(),
		config: &config.ConfigureApp,
		db:     sqlx.NewDb(db, "sqlmock"),
	}

	mock.ExpectQuery("SELECT (.+) FROM users").WithArgs().WillReturnRows(&sqlmock.Rows{})

	result, err := repo.FindOneUser(entity.FindOneUser{
		Username: new(string),
		Email:    new(string),
	}, []string{})

	assert.NoError(t, err, "Unexpected error")
	assert.Nil(t, result, "Unexpected result")
	assert.NoError(t, mock.ExpectationsWereMet(), "Expectations not met")
}
