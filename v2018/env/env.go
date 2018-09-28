package env

import "os"

// GetEnvOrElse return the current value of env variable
// or defaultValue
func GetEnvOrElse(env string, defaultValue string) string {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		return v
	default:
		return defaultValue
	}
}
