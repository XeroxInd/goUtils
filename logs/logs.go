package logs

import (
	"log"

	"github.com/getsentry/raven-go"
)

func Info(msg string, category string) {
	raven.CaptureMessage(msg, map[string]string{"category": category})
	log.Print(msg)
}

func Error(msg error, category string) {
	raven.CaptureError(msg, map[string]string{"category": category})
	log.Print(msg)
}
