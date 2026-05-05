package repository

import (
	"finance-api/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(infoUser *model.User) error {
	query := `INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, infoUser.Name, infoUser.Email, infoUser.Password, infoUser.CreatedAt)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.Get(&user, query, email)
	return &user, err
}
