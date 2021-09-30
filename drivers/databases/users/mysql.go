package users

import (
	"context"
	"cookly/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (repository *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user Users
	result := repository.Conn.First(&user, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (repository *MysqlUserRepository) Register(ctx context.Context, userDomain *users.Domain) error {
	reg := FromDomain(*userDomain)

	result := repository.Conn.Create(&reg)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlUserRepository) GetUserByID(ctx context.Context, id int) (users.Domain, error) {
	getUserByID := Users{}

	result := repository.Conn.First(&getUserByID, id)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return getUserByID.ToDomain(), nil
}
