/**
 * Package config
 * @file      : config.go
 * @author    : xaoyaoyao
 * @version   : 1.0.0
 * @time      : 2025/2/18 10:00
 **/

package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var cfg Config

func Init() error {
	err := InitBaseConfig()
	if err != nil {
		fmt.Println("Error initializing base config:", err)
		return err
	}
	return nil
}

type Config struct {
	Addr     string // 服务地址IP
	LogLevel string // 日志级别

	ApiKey    string // API Key
	SecretKey string // API Secret

	SessionExpiresIn      int // Session会话有效时间，单位秒
	AccessTokenExpiresIn  int // 访问令牌的有效时间，单位秒
	RefreshTokenExpiresIn int // 刷新令牌的有效时间，单位秒

	Issuer string // Issuer
}

func Get() Config {
	return cfg
}

func IsLocalEnv() bool {
	envType := os.Getenv("RUN_ENV")
	return envType == "local"
}

func InitBaseConfig() error {
	//if IsLocalEnv() {
	filePath := ".env"
	err := godotenv.Load(filePath)
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
	return loadBaseConfigByEnv()
	//}
	//return loadBaseConfig()
}

func loadBaseConfigByEnv() error {
	cfg = Config{
		Addr:                  os.Getenv("ADDR"),
		LogLevel:              os.Getenv("LOG_LEVEL"),
		ApiKey:                os.Getenv("API_KEY"),
		SecretKey:             os.Getenv("SECRET_KEY"),
		SessionExpiresIn:      setValueToInt(os.Getenv("SESSION_EXPIRES_IN"), 3600),
		AccessTokenExpiresIn:  setValueToInt(os.Getenv("ACCESS_TOKEN_EXPIRES_IN"), 2592000),
		RefreshTokenExpiresIn: setValueToInt(os.Getenv("REFRESH_TOKEN_EXPIRES_IN"), 7776000),

		Issuer: os.Getenv("ISSUER"),
	}
	return nil
}

func loadBaseConfig() error {
	//parameters, err := LoadParameterFile(ssmParameterFile)
	//if parameters == nil {
	//	err = fmt.Errorf("error loading ssm parameter")
	//	panic(err)
	//}
	//if err != nil {
	//	log.Fatal("Error loading ssm parameter", err)
	//	return err
	//}
	//err = json.Unmarshal([]byte(*parameters), &cfg)
	//if err != nil {
	//	log.Fatal("Error loading ssm parameter", err)
	//	return err
	//}
	log.Println("loading parameter success...")
	return nil
}

func setValueToInt(val string, defaultValue int) int {
	if val == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return intValue
}
