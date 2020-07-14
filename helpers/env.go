package helpers

import "os"

func EnvValReplace(value *string , key string) {
	if envVal := os.Getenv(key); envVal != "" {
		*value = envVal
	}
}

func EnvGetByDefault(key,defVal string) string {
	value := defVal
	EnvValReplace(&value,key)
	return value
}