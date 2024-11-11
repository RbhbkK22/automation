package repositories

import (
	"automation/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetUserByLogin(login string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	query := `SELECT w.id, w.login, w.fio, p.name AS post, w.pass
FROM workers w
JOIN positions p ON w.post = p.id
WHERE w.login =  ?`
	row := r.db.QueryRow(query, login)
	err := row.Scan(&user.ID, &user.Login, &user.Fio, &user.Post, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	return &user, nil
}
