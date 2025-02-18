/**
 * Package i18n
 * @file      : i18n.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:10
 **/

package i18n

import (
	"embed"
	gini18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

//go:embed localize/*
var FS embed.FS

// GetMessage get the i18n push without error handling
// code is one of these type: error code
// Example:
// GetMessage(context, 200) // code is 200
func GetMessage(c *gin.Context, code int) string {
	codeStr := strconv.Itoa(code) + "_code"
	message := gini18n.MustGetMessage(c, codeStr)
	return trimSuffix(message)
}

// MustGetMessage get the i18n push without error handling
// param is one of these type: messageID, *i18n.LocalizeConfig
// Example:
// MustGetMessage(context, "hello") // messageID is hello
func MustGetMessage(c *gin.Context, messageID string) string {
	message := gini18n.MustGetMessage(c, messageID)
	return trimSuffix(message)
}

func trimSuffix(message string) string {
	if message != "" && strings.HasSuffix(message, "\n") {
		message = strings.TrimSuffix(message, "\n")
	}
	return message
}
