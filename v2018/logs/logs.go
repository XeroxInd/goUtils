package logs

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/getsentry/raven-go"
)

// Info msg to std out and send an info event to the Sentry server
// pointed by `SENTRY_DSN` environment variable.
func Info(msg string, category string) {
	raven.CaptureMessage(msg, map[string]string{"category": category})
	log.Print(msg)
}

// Error msg to std out and send an error event to the Sentry server
// pointed by `SENTRY_DSN` environment variable.
func Error(msg error, category string) {
	raven.CaptureError(msg, map[string]string{"category": category})
	log.Print(msg)
}

func PrettyPrint(input interface{}) (string, error) {
	var log = new(bytes.Buffer)
	encoder := json.NewEncoder(log)
	encoder.SetIndent("", "   ")
	err := encoder.Encode(input)
	if err != nil {
		return "", err
	}
	return log.String(), nil
}
