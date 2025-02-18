/**
 * Package dto
 * @file      : claim.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 15:49
 **/

package dto

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
