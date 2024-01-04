package env

import "os"

func GetEnv(key string, defValue string) string {
	envData := os.Getenv(key)

	if envData == "" {
		envData = defValue
	}

	return envData
}
