package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Roongkun/software-eng-ii/internal/third-party/databases"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

type JwtClaim struct {
	Email   string
	IsAdmin bool
	jwt.RegisteredClaims
}

func (j *JwtWrapper) GenerateToken(ctx context.Context, email string, isAdmin bool) (signedToken string, err error) {
	claims := &JwtClaim{
		Email:   email,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(j.ExpirationMinutes))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	if err := databases.RedisClient.Set(ctx, signedToken, email, 60*time.Minute).Err(); err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JwtWrapper) ValidateToken(c *gin.Context, signedToken string, isAdmin bool) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JwtClaim{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		c.Set("errorStatus", http.StatusUnauthorized)
		c.Set("errorMessage", err.Error())
		return nil, err
	} else if claims, ok := token.Claims.(*JwtClaim); ok {
		if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
			c.Set("errorStatus", http.StatusUnauthorized)
			c.Set("errorMessage", "the session has expired, token refreshing is needed")
			return nil, errors.New("the session has expired, token refreshing is needed")
		} else {
			if isAdmin && !claims.IsAdmin {
				c.Set("errorStatus", http.StatusBadRequest)
				c.Set("errorMessage", "the email provided is not an administrator email")
				return nil, errors.New("the email provided is not an administrator email")
			}
			if !isAdmin && claims.IsAdmin {
				c.Set("errorStatus", http.StatusBadRequest)
				c.Set("errorMessage", "the email provided is an administrator email, please use the /admin path instead")
				return nil, errors.New("the email provided is an administrator email, please use the /admin path instead")
			}

			return claims, nil
		}
	} else {
		c.Set("errorStatus", http.StatusInternalServerError)
		c.Set("errorMessage", "unknown claim type, cannot proceed")
		return nil, errors.New("unknown claim type, cannot proceed")
	}
}
