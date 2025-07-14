package usecase

import (
	"cleanarchitecture/internal/domain"
	"cleanarchitecture/internal/middleware"
	"cleanarchitecture/internal/repository"
)

type UserUsecase interface {
	Authenticate(username, password string) (bool, error)
	GenerateToken(username string) (string, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Authenticate(username, password string) (bool, error) {
	user, err := u.repo.FindByUsername(username)
	if err != nil {
		return false, err
	}
	if user.Password != password {
		return false, domain.ErrInvalidCredentials
	}
	return true, nil
}

func (u *userUsecase) GenerateToken(username string) (string, error) {
	user, err := u.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateToken(*user)

	if err != nil {
		return "", err
	}
	return token, nil
}
