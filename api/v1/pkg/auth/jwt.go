package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/topboyasante/mrkt-api/internal/config"
)

func CreateJWTTokens(data any) (string, string, int64, error) {
	accessTokenExp := time.Now().Add(time.Hour * 24 * 15).Unix()
	accessTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": data,              // Subject (user identifier)
		"iss": "mrkt",            // Issuer
		"aud": data,              // Audience (user role)
		"exp": accessTokenExp,    // Expiration time
		"iat": time.Now().Unix(), // Issued at
	})

	accessTokenString, err := accessTokenClaims.SignedString([]byte(config.ENV.JWTKey))
	if err != nil {
		return "", "", 0, err
	}

	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": data,                                       // Subject (user identifier)
		"iss": "mrkt",                                     // Issuer
		"aud": data,                                       // Audience (user role)
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Expiration time = 30 days
		"iat": time.Now().Unix(),                          // Issued at
	})

	refreshTokenString, err := refreshTokenClaims.SignedString([]byte(config.ENV.JWTKey))
	if err != nil {
		return "", "", 0, err
	}

	return accessTokenString, refreshTokenString, accessTokenExp, nil
}
