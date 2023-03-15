package service

import (
	"context"
	"errors"

	"github.com/rarifbkn/api-rest-gorm-echo/encryption"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/models"
)

var (
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvallidCredentials = errors.New("Invalid credentials")
	ErrRoleAlreadyAdded    = errors.New("Role was already added for this user")
	ErrRoleNotFound        = errors.New("Role not found ")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	//Encryption password
	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(bb) //contraseña cifrada

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	//Realiza la desencriptacion de la contraseña del usuario

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptPassword) != password {
		return nil, ErrInvallidCredentials
	}
	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) AddUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return nil
	}

	for _, r := range roles {
		if r.RoleID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRoles(ctx, userID)
	if err != nil {
		return nil
	}

	roleFound := false
	for _, r := range roles {
		if r.RoleID == roleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}
	return s.repo.RemoveUserRole(ctx, userID, roleID)
}
