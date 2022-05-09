package dto

import (
	"time"
	"toko/cmd/domain/user/entity"
	"toko/pkg/auth/dto"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserAuthResponse struct {
	Type         string `json:"type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserListResponse []*UserResponse

func CreateUserResponse(user *entity.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func CreateUserListResponse(users *entity.UserList) UserListResponse {
	userResp := UserListResponse{}
	for _, p := range *users {
		user := CreateUserResponse(p)
		userResp = append(userResp, &user)
	}
	return userResp
}

func CreateUserAuthResponse(token dto.AccessToken) UserAuthResponse {
	return UserAuthResponse{
		Type:         token.Type,
		Token:        token.Token,
		RefreshToken: token.RefreshToken,
	}
}
