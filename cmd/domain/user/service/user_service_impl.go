package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
	"toko/cmd/domain/user/dto"
	"toko/cmd/domain/user/entity"
	"toko/cmd/domain/user/repository"
	"toko/internal/protocol/http/errors"
	"toko/pkg/auth"
	"toko/pkg/hash"
)

type UserServiceImpl struct {
	RepoUser repository.UserRepository
	JwtAuth  auth.JwtToken
	repoOnce sync.Once
}

func (s UserServiceImpl) GetUsers() (*dto.UserListResponse, error) {
	users, err := s.RepoUser.FindAll()
	if err != nil {
		log.Err(err).Msg("Error fetch users from DB")
		return nil, err
	}
	usersResp := dto.CreateUserListResponse(users)
	return &usersResp, nil
}

func (s UserServiceImpl) GetUserById(userId uint) (*dto.UserResponse, error) {
	user, err := s.RepoUser.Find(userId)
	if err != nil {
		log.Err(err).Msg("Error fetch user from DB")
		return nil, err
	}
	userResp := dto.CreateUserResponse(user)
	return &userResp, nil
}

func (s UserServiceImpl) Store(request *dto.UserRequestBody) (*dto.UserResponse, error) {
	passwordHashed, err := hash.AppBcryptImpl{}.HashAndSalt([]byte(request.Password))
	if err != nil {
		log.Err(err).Msg("Error hash password to bcrypt")
		return nil, err
	}

	userRepo, err := s.RepoUser.Insert(&entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: passwordHashed,
	})

	if err != nil {
		log.Err(err).Msg("Error insert user to DB")
		return nil, err
	}

	userResp := dto.CreateUserResponse(userRepo)
	log.Err(err).Msg("Successfully insert to to DB")
	return &userResp, nil
}

func (s UserServiceImpl) Login(request *dto.UserRequestLogin) (*dto.UserAuthResponse, error) {
	user, err := s.RepoUser.FindByEmail(request.Email)
	if err != nil {
		log.Err(err).Msg("Error fetch user from DB")
		return nil, errors.FindErrorType(err)
	}

	isMatched := hash.AppBcryptImpl{}.ComparePasswords(user.Password, []byte(request.Password))

	if !isMatched {
		log.Err(err).Msg("email and password didn't match")
		return nil, errors.Unauthorization("email and password didn't match")
	}

	accessToken := s.JwtAuth.Sign(jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	authResp := dto.CreateUserAuthResponse(accessToken)

	return &authResp, nil
}

func (s UserServiceImpl) Refresh(userId uint) (*dto.UserAuthResponse, error) {
	user, err := s.RepoUser.Find(userId)
	if err != nil {
		log.Err(err).Msg("Error fetch user from DB")
		return nil, errors.FindErrorType(err)
	}

	accessToken := s.JwtAuth.Sign(jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	authResp := dto.CreateUserAuthResponse(accessToken)

	return &authResp, nil
}
