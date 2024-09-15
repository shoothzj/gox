package envx

import (
	"os"
	"strconv"
)

func GetEnvStr(key string, value string) string {
	aux := os.Getenv(key)
	if aux != "" {
		return aux
	}
	return value
}

func GetEnvInt(key string, value int) int {
	aux := os.Getenv(key)
	if aux == "" {
		return value
	}
	res, err := strconv.Atoi(aux)
	if err != nil {
		return value
	}
	return res
}

func GetEnvInt64(key string, value int64) int64 {
	aux := os.Getenv(key)
	if aux == "" {
		return value
	}
	res, err := strconv.ParseInt(aux, 10, 64)
	if err != nil {
		return value
	}
	return res
}

func GetEnvBool(key string, value bool) bool {
	aux := os.Getenv(key)
	if aux == "" {
		return value
	}
	return aux == "true"
}
