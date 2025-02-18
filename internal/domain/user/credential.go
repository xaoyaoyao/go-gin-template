/**
 * Package user
 * @file      : credential.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 16:11
 **/

package user

import (
	"encoding/gob"
	"fmt"
	"github.com/coverai/api/internal/common/util"
	"github.com/coverai/api/internal/config"
	"github.com/coverai/api/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CredentialResponse struct {
	AccessToken           string    `json:"access_token"`
	ExpiresIn             int       `json:"expires_in"`
	UserId                string    `json:"user_id"`
	RefreshToken          string    `json:"refresh_token"`
	TokenType             string    `json:"token_type"`
	Scope                 string    `json:"scope,omitempty"`
	ExpiresAt             time.Time `json:"expires_at"`
	RefreshTokenExpiresIn int       `json:"refresh_token_expires_in"`
}

type Credential struct {
}

func init() {
	gob.Register(&CredentialResponse{})
}

func NewCredential() *Credential {
	return &Credential{}
}

func (credential Credential) Auth(c *gin.Context, userId string) (*CredentialResponse, error) {
	secretKey := config.Get().ApiKey + ":" + config.Get().SecretKey
	expiresIn := config.Get().AccessTokenExpiresIn
	accessToken, err := credential.GenerateToken(userId, secretKey, expiresIn)
	if err != nil {
		return nil, err
	}

	refreshTokenExpiresIn := config.Get().RefreshTokenExpiresIn
	refreshToken, err := credential.GenerateToken(userId, secretKey, refreshTokenExpiresIn)
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().Add(time.Second * time.Duration(expiresIn))
	return &CredentialResponse{
		AccessToken:           accessToken,
		ExpiresIn:             expiresIn,
		UserId:                userId,
		RefreshToken:          refreshToken,
		TokenType:             "Bearer",
		Scope:                 "scope",
		ExpiresAt:             expiresAt,
		RefreshTokenExpiresIn: refreshTokenExpiresIn,
	}, nil
}

func (credential Credential) GenerateToken(userId, secretKey string, expiresAt int) (string, error) {
	tokenExpiresAt := time.Duration(expiresAt) * time.Second
	claims := &dto.Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiresAt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.Get().Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func (credential Credential) CheckToken(c *gin.Context, token string) (*dto.Claims, error) {
	secretKey := config.Get().ApiKey + ":" + config.Get().SecretKey
	claims, err := credential.ParseToken(token, secretKey)
	if err != nil {
		return nil, err
	}
	c.Set(util.IdKey, claims.UserID)
	c.Set(util.ClaimKey, *claims)
	return claims, nil
}

func (credential Credential) ParseToken(tokenString string, secretKey string) (*dto.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
