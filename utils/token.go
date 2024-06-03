package utils

import (
	"fmt"
	"strconv"
	"time"

	"srating/domain"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	ID := strconv.Itoa(int(user.ID))
	Role := user.Role
	claims := &domain.JwtCustomClaims{
		Role: Role,
		ID:   ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	ID := strconv.Itoa(int(user.ID))
	Role := user.Role
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID:   ID,
		Role: Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	return token.SignedString([]byte(secret))
}

func ParseJWTToken(requestToken, secret string) (*jwt.Token, error) {
	return jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func IsAuthorized(requestToken, secret string) (bool, error) {
	_, err := ParseJWTToken(requestToken, secret)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken, secret string) (string, error) {
	token, err := ParseJWTToken(requestToken, secret)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}
	return claims["id"].(string), nil
}

func ExtractRoleFromToken(requestToken, secret string) (string, error) {
	token, err := ParseJWTToken(requestToken, secret)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["role"].(string), nil
}

func ExtractLocationFromToken(requestToken, secret string) (string, error) {
	token, err := ParseJWTToken(requestToken, secret)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["location"].(string), nil
}
