/**
 * Package util
 * @file      : tool.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:19
 **/

package util

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

func NewId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func SetTimestamp(t time.Time) int64 {
	timestamp := t.UnixNano()
	millisecondTimestamp := timestamp / 1e6
	return millisecondTimestamp
}

// GetPriority get priority by time to millisecond timestamp
func GetPriority(t time.Time) int64 {
	timestamp := t.UnixNano()
	millisecondTimestamp := timestamp / 1e6
	return millisecondTimestamp
}

// ExtractToken Get Authorization Value
func ExtractToken(authorizationValue string) string {
	if authorizationValue == "" {
		return authorizationValue
	}
	if strings.HasPrefix(authorizationValue, BearerKey) {
		return strings.TrimPrefix(authorizationValue, BearerKey)
	}
	return authorizationValue
}
