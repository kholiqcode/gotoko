package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"toko/config"
	"toko/pkg/auth/dto"
	"time"
)

type JwtToken interface {
	Sign(claims jwt.MapClaims) dto.AccessToken
}

type JwtTokenImpl struct {
	jwtTokenTimeExp        time.Duration
	jwtRefreshTokenTimeExp time.Duration
}

func NewJwtToken() *JwtTokenImpl {
	jwtTokenDuration, err := time.ParseDuration(config.Get().JwtAccessExpired)

	if err != nil {
		log.Err(err).Msg(config.Get().JwtAccessExpired)
	}
	jwtRefreshDuration, err := time.ParseDuration(config.Get().JwtRefreshExpired)
	if err != nil {
		log.Err(err).Msg(config.Get().JwtRefreshExpired)
	}
	return &JwtTokenImpl{
		jwtTokenTimeExp:        jwtTokenDuration,
		jwtRefreshTokenTimeExp: jwtRefreshDuration,
	}
}

func (o JwtTokenImpl) Sign(claims jwt.MapClaims) dto.AccessToken {
	timeNow := time.Now()
	tokenExpired := timeNow.Add(o.jwtTokenTimeExp).Unix()

	if claims["id"] == nil {
		return dto.AccessToken{}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// setup userdata
	var _, checkExp = claims["exp"]
	var _, checkIat = claims["iat"]

	// if user didn't define claims expired
	if !checkExp {
		claims["exp"] = tokenExpired
	}
	// if user didn't define claims iat
	if !checkIat {
		claims["iat"] = timeNow.Unix()
	}
	claims["token_type"] = "access_token"

	token.Claims = claims

	authToken := new(dto.AccessToken)
	tokenString, err := token.SignedString([]byte(config.Get().JwtSecretKey))

	if err != nil {
		log.Err(err).Msg("Can't create access token")
		return dto.AccessToken{}
	}

	authToken.Token = tokenString
	authToken.Type = "Bearer"

	//create refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenExpired := timeNow.Add(o.jwtRefreshTokenTimeExp).Unix()

	claims["exp"] = refreshTokenExpired
	claims["token_type"] = "refresh_token"
	refreshToken.Claims = claims

	refreshTokenString, err := refreshToken.SignedString([]byte(config.Get().JwtSecretKey))

	if err != nil {
		return dto.AccessToken{}
	}
	authToken.RefreshToken = refreshTokenString

	return dto.AccessToken{
		Type:         authToken.Type,
		Token:        authToken.Token,
		RefreshToken: authToken.RefreshToken,
	}
}
