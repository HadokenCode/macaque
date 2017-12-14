package env

import "os"

//GetEnv  It returns env var or default value for a key
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
