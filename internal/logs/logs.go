/**
 * Package logs
 * @file      : logs.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 09:58
 **/

package logs

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	LoggerKey = "Logger"
	SessionId = "SessionId"
	TraceId   = "TraceId"
)

func Init(logLevel string) {
	logrus.StandardLogger()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stdout)
	if gin.Mode() == gin.ReleaseMode {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: time.RFC3339Nano,
		})
	}
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	logrus.SetReportCaller(true)
	if level > logrus.DebugLevel {
		logrus.SetReportCaller(true)
	}
}

func FromContext(ctx context.Context) *logrus.Entry {
	if log, ok := ctx.Value(LoggerKey).(*logrus.Entry); ok {
		return log
	}
	return logrus.NewEntry(logrus.StandardLogger())
}

func ToContext(ctx context.Context, log *logrus.Entry) context.Context {
	return context.WithValue(ctx, LoggerKey, log)
}
