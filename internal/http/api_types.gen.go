// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package http

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// BasicQueryParams defines model for BasicQueryParams.
type BasicQueryParams struct {
	// Language Language, zh_cn, en, default en
	Language *string `json:"language,omitempty"`

	// Os Device version information
	Os *string `json:"os,omitempty"`

	// Sid Session ID
	Sid *string `json:"sid,omitempty"`

	// Version The app's version
	Version *string `json:"version,omitempty"`
}

// CredentialVO defines model for CredentialVO.
type CredentialVO struct {
	// AccessToken access token
	AccessToken string `json:"accessToken"`

	// ExpiresIn The validity period of the access token, in seconds
	ExpiresIn int `json:"expiresIn"`

	// Id user id
	Id string `json:"id"`

	// RefreshToken refresh token
	RefreshToken string `json:"refreshToken"`

	// RefreshTokenExpiresIn The validity period of the refresh token, in seconds
	RefreshTokenExpiresIn int `json:"refreshTokenExpiresIn"`

	// Scope Authorization scope of access token
	Scope *string `json:"scope,omitempty"`

	// TokenType Token type, usually bearer
	TokenType *string `json:"tokenType,omitempty"`
}

// InitializationDataVO defines model for InitializationDataVO.
type InitializationDataVO map[string]interface{}

// ResponseEntity defines model for ResponseEntity.
type ResponseEntity struct {
	// Code Response code.
	Code int32 `json:"code"`

	// Msg Message about the response.
	Msg string `json:"msg"`
}

// InitializationDataParams defines parameters for InitializationData.
type InitializationDataParams struct {
	QueryParams *BasicQueryParams `form:"queryParams,omitempty" json:"queryParams,omitempty"`
}

// SignupJSONBody defines parameters for Signup.
type SignupJSONBody struct {
	// DeviceId the device's id
	DeviceId string `json:"deviceId"`
}

// SignupParams defines parameters for Signup.
type SignupParams struct {
	// GOAPIKEY API KEY
	GOAPIKEY string `json:"GO-API-KEY"`
}

// SignupJSONRequestBody defines body for Signup for application/json ContentType.
type SignupJSONRequestBody SignupJSONBody
