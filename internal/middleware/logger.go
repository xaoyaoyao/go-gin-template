/**
 * Package middleware
 * @file      : logger.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:16
 **/

package middleware

import (
	"bytes"
	"fmt"
	"github.com/coverai/api/internal/common/util"
	"github.com/coverai/api/internal/logs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"time"
)

func extractTraceId(c *gin.Context) string {
	// Get trace id from header
	if c.GetHeader(logs.TraceId) != "" {
		traceId := c.GetHeader(logs.TraceId)
		return traceId
	}
	b, err := c.GetRawData()
	if err != nil {
		logs.FromContext(c).Error(err)
		return ""
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
	timestamp := util.GetPriority(time.Now())
	return strconv.FormatInt(timestamp, 10)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		reqTraceId := extractTraceId(c)

		logEntry := logs.FromContext(c).WithFields(logrus.Fields{
			logs.TraceId: reqTraceId,
		})

		c.Set(string(logs.TraceId), reqTraceId)
		c.Set(string(logs.LoggerKey), logEntry)

		// Process request
		c.Next()
		// After request

		// Combine request path
		path := c.Request.URL.Path
		if raw := c.Request.URL.RawQuery; raw != "" {
			path = path + "?" + raw
		}

		// Response status code
		statusCode := c.Writer.Status()

		logEntry = logs.FromContext(c).WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"status":  statusCode,
			"path":    path,
			"latency": fmt.Sprintf("%dµs", time.Since(t).Microseconds()),
		})

		switch {
		case statusCode >= 400 && statusCode < 500:
			logEntry.Warn()
		case statusCode >= 500:
			logEntry.Error(c.Errors.Errors())
		default:
			logEntry.Info()
		}
	}
}
