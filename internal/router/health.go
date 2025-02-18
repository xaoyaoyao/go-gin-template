/**
 * Package router
 * @file      : health.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:00
 **/

package router

import (
	"github.com/coverai/api/internal/common/rsp"
	"github.com/coverai/api/internal/config"
	"github.com/coverai/api/internal/i18n"
	"github.com/coverai/api/internal/logs"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
)

type Health struct {
	Code int `yaml:"code"`
}

func getHealthFile() string {
	filePath := "/app/health/health.yaml"
	if config.IsLocalEnv() {
		filePath = "./health/health.yaml"
	}
	return filePath
}

func HealthHandler(c *gin.Context) {
	// Read configuration file status code
	filePath := getHealthFile()
	code, err := readStatusCodeFromFile(c, filePath)
	if err != nil {
		logs.FromContext(c).Error(err)
		code = http.StatusOK
	}
	healthCheckHandler(c, code)
}

func HealthCheck(c *gin.Context) {
	healthCheckHandler(c, http.StatusOK)
}

func healthCheckHandler(c *gin.Context, code int) {
	rs := rsp.ResponseEntity{
		Code:    int32(code),
		Message: i18n.GetMessage(c, code),
	}
	c.JSON(code, rs)
}

func readStatusCodeFromFile(c *gin.Context, filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logs.FromContext(c).Error(err)
		return http.StatusOK, nil
	}
	defer file.Close()
	var health Health
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&health)
	if err != nil {
		logs.FromContext(c).Error(err)
		return http.StatusOK, nil
	}
	return health.Code, nil
}

func UpdateHealthHandler(c *gin.Context) {
	status := http.StatusOK
	code := c.Param("code")
	if code == "" || code == "503" {
		status = http.StatusServiceUnavailable
	}
	filePath := getHealthFile()
	newHealth := Health{
		Code: status,
	}
	data, err := yaml.Marshal(newHealth)
	if err != nil {
		logs.FromContext(c).Error(err)
		return
	}
	_ = os.WriteFile(filePath, data, 0644)
}
