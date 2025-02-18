/**
 * Package util
 * @file      : constant.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:28
 **/

package util

var (
	RootPath     = "/api"
	VersionPath  = "/v1"
	Version2Path = "/v2"

	IdKey    = "user_id"
	ClaimKey = "userClaims"
)

const (
	SessionIdKey          = "sessionId"
	QuerySessionIdKey     = "sid"
	HeaderSessionIdKey    = "X-SESSION-ID"
	AuthorizationKey      = "Authorization"
	BearerKey             = "Bearer "
	SessionKeyPrefix      = ""
	AccessTokenKeyPrefix  = "ACCESS_"
	RefreshTokenKeyPrefix = "REFRESH_"
	UserTokenKeyPrefix    = "USER_"
	RecordingsTimePeriod  = "X-Time-Period"
)
