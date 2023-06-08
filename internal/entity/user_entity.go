package entity

import "time"

type FindOneUser struct {
	ID        *int       `json:"id" db:"id"`
	Username  *string    `json:"username" db:"username"`
	Email     *string    `json:"email" db:"email"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type Users struct {
	ID        *int       `json:"id" db:"id"`
	Username  *string    `json:"username" db:"username"`
	Email     *string    `json:"email" db:"email"`
	Password  *string    `json:"password" db:"password"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type CreateUser struct {
	Username  *string    `json:"username" db:"username"`
	Email     *string    `json:"email" db:"email"`
	Password  *string    `json:"password" db:"password"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
