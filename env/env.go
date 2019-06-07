package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func GetEnvOrElse(env string, defaultValue string) string {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		return v
	default:
		return defaultValue
	}
}

func GetIntEnvOrElse(env string, defaultValue int) (value int, parseError error) {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		val, err := strconv.Atoi(v)
		if err != nil {
			return defaultValue, errors.Wrap(err, fmt.Sprintf("error when trying to parse int for %s", env))
		}
		return val, err
	default:
		return defaultValue, err
	}
}

func GetFloatEnvOrElse(env string, defaultValue float64) (value float64, parseError error) {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defaultValue, errors.Wrap(err, fmt.Sprintf("error when trying to parse flot64 for %s", env))
		}
		return val, err
	default:
		return defaultValue, err
	}
}

func GetBoolEnvOrElse(env string, defaultValue bool) (value bool, parseError error) {
	v, ok := os.LookupEnv(env)
	switch ok {
	case true:
		val, err := strconv.ParseBool(v)
		if err != nil {
			return defaultValue, errors.Wrap(err, fmt.Sprintf("error when trying to parse bool for %s", env))
		}
		return val, err
	default:
		return defaultValue, err
	}
}
