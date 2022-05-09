package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"toko/config"
	"toko/internal/protocol/http/response"
	"time"
)

func JwtVerifyAccess() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Get().JwtSecretKey),
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				claims := t.Claims.(jwt.MapClaims)
				rawExp := claims["exp"].(float64)
				exp := int64(rawExp)
				if t.Method.Alg() != "HS256" {
					log.Err(errors.New("unexpected signing method")).Msgf("unexpected signing method: %v", t.Method.Alg())
					return nil, errors.New("unexpected jwt signing method")
				} else if claims["token_type"] != "access_token" {
					log.Err(errors.New("unexpected token type")).Msgf("unexpected token type: %v", claims["token_type"])
					response.Err(c, fmt.Errorf("unexpected token type %v", claims["token_type"]))
					return nil, fmt.Errorf("unexpected token type %v", claims["token_type"])
				} else if exp < time.Now().Unix() {
					log.Err(errors.New("token expired")).Msgf("token expired: %v", exp)
					response.Err(c, errors.New("token expired"))
					return nil, errors.New("token expired")
				}

				return []byte(config.Get().JwtSecretKey), nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				log.Err(errors.New("invalid token")).Msgf("invalid token: %v", token.Raw)
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	})
}

func JwtVerifyRefresh() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.Get().JwtSecretKey),
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				claims := t.Claims.(jwt.MapClaims)
				rawExp := claims["exp"].(float64)
				exp := int64(rawExp)
				if t.Method.Alg() != "HS256" {
					log.Err(errors.New("unexpected signing method")).Msgf("unexpected signing method: %v", t.Method.Alg())
					return nil, errors.New("unexpected jwt signing method")
				} else if claims["token_type"] != "refresh_token" {
					log.Err(errors.New("unexpected token type")).Msgf("unexpected token type: %v", claims["token_type"])
					response.Err(c, fmt.Errorf("unexpected token type %v", claims["token_type"]))
					return nil, fmt.Errorf("unexpected token type %v", claims["token_type"])
				} else if exp < time.Now().Unix() {
					log.Err(errors.New("token expired")).Msgf("token expired: %v", exp)
					response.Err(c, errors.New("token expired"))
					return nil, errors.New("token expired")
				}

				return []byte(config.Get().JwtSecretKey), nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				log.Err(errors.New("invalid token")).Msgf("invalid token: %v", token.Raw)
				return nil, errors.New("invalid token")
			}
			//Send User ID To Context
			c.Set("user_id", token.Claims.(jwt.MapClaims)["id"])

			return token, nil
		},
	})
}

//func JwtVerifyRefresh(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		if err := next(c); err != nil {
//			c.Error(err)
//		}
//		user := c.Get("user").(*jwt.Token)
//		claims := user.Claims.(jwt.MapClaims)
//		tokenType := claims["token_type"].(string)
//		id := claims["id"].(string)
//		rawExp := claims["exp"].(float64)
//		exp := int64(rawExp)
//
//		if _, ok := user.Method.(*jwt.SigningMethodRSA); !ok {
//			log.Err(errors.New("unexpected signing method")).Msgf("unexpected signing method: %v", user.Header["alg"])
//			response.Err(c, errors.New("unexpected signing method"))
//			return errors.New("unexpected signing method")
//		} else if tokenType != "refresh_token" {
//			log.Err(errors.New("unexpected token type")).Msgf("unexpected signing method: %v", tokenType)
//			response.Err(c, errors.New("unexpected token type"))
//			return errors.New("unexpected token type")
//		} else if !user.Valid {
//			log.Err(errors.New("token is not valid")).Msgf("token is not valid: %v", user.Raw)
//			response.Err(c, errors.New("token is not valid"))
//			return errors.New("token is not valid")
//		} else if id == "" {
//			log.Err(errors.New("user not found")).Msgf("user not found: %v", id)
//			response.Err(c, errors.New("user not found"))
//			return errors.New("user not found")
//		} else if exp < time.Now().Unix() {
//			log.Err(errors.New("token expired")).Msgf("token expired: %v", id)
//			response.Err(c, errors.New("token expired"))
//			return errors.New("token expired")
//		}
//		return next(c)
//	}
//}
