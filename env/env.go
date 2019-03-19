package env

import (
	"os"
	"strconv"
)

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

func GetIntEnvOrElse(env string, defaultValue int) int {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		val, err := strconv.Atoi(v)
		if err != nil {
			return defaultValue
		}
		return val
	default:
		return defaultValue
	}
}

func GetFloatEnvOrElse(env string, defaultValue float64) float64 {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defaultValue
		}
		return val
	default:
		return defaultValue
	}
}

func GetBoolEnvOrElse(env string, defaultValue bool) bool {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		val, err := strconv.ParseBool(v)
		if err != nil {
			return defaultValue
		}
		return val
	default:
		return defaultValue
	}
}
