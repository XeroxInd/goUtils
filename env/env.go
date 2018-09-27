package env

import "os"

func GetEnvOrElse(env string, defaultValue string) string {
	v, ok := os.LookupEnv(env)
	if ok {
		return v
	} else {
		return defaultValue
	}
}
