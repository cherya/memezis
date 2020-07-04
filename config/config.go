package config

import (
	"fmt"
	"os"
	"strconv"
)

type configKey string

const (
	HttpPort      configKey = "HTTP_PORT"
	GrpcPort      configKey = "GRPC_PORT"
	DatabaseDsn   configKey = "DATABASE_DSN"
	RedisAddress  configKey = "REDIS_ADDRESS"
	RedisPassword configKey = "REDIS_PASSWORD"
	S3BucketURL   configKey = "S3_BUCKET_URL"
	S3Key         configKey = "S3_KEY"
	S3Secret      configKey = "S3_SECRET"
	S3Region      configKey = "S3_REGION"
	S3Endpoint    configKey = "S3_ENDPOINT"
	S3BucketName  configKey = "S3_BUCKET_NAME"
	BotClientKey  configKey = "BOT_CLIENT_KEY"
)

func GetValue(key configKey) string {
	return os.Getenv(string(key))
}

func GetInt(key configKey) int {
	val, err := strconv.Atoi(os.Getenv(string(key)))
	if err != nil {
		panic(fmt.Sprintf("%s env value is not integer", string(key)))
	}
	return val
}

func GetInt64(key configKey) int64 {
	val := GetInt(key)
	return int64(val)
}

func GetBool(key configKey) bool {
	val, found := os.LookupEnv(string(key))
	if val == "false" {
		return false
	}
	return found
}
