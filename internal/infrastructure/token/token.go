package token

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService struct {
	tokenTTL   time.Duration
	signingKey string
}

func NewTokenService(tokenTTL int, signingKey string) *TokenService {
	return &TokenService{
		tokenTTL:   time.Duration(tokenTTL) * time.Hour,
		signingKey: signingKey,
	}
}

type UserJWTClaims struct {
	jwt.StandardClaims
	UserId   int `json:"user_id"`
	UserRole int `json:"user_role"`
}

func (s *TokenService) GenerateToken(ctx context.Context, id, role int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserJWTClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
		role,
	})

	signedToken, err := token.SignedString([]byte(s.signingKey))
	if err != nil {
		return "", fmt.Errorf("TokenService - SignIn - token.SignedString: %w", err)
	}

	return signedToken, nil
}

func (s *TokenService) ParseToken(ctx context.Context, tokenString string) (int, int, error) {
	t, err := jwt.ParseWithClaims(tokenString, &UserJWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(s.signingKey), nil
	})

	if err != nil {
		return 0, 0, fmt.Errorf("TokenService - ParseToken: %w", err)
	}

	claims, ok := t.Claims.(*UserJWTClaims)
	if !ok {
		return 0, 0, fmt.Errorf("invalid claims type")
	}

	return claims.UserId, claims.UserRole, nil
}
