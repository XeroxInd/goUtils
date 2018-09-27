package logs

import (
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
