package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/getsentry/raven-go"
)

// Info msg to std out and send an info event to the Sentry server
// pointed by `SENTRY_DSN` environment variable.
// tags can contain a map of additional tags information
// /!\ This is a non blocking call, if the program die just after the call, the message can be lost
func InfoWithTags(msg string, tags map[string]string) {
	packet := &raven.Packet{
		Message: msg,
		Level:   raven.INFO,
	}
	raven.Capture(packet, tags)
	log.Printf("%s, %s", "[INFO]", msg)
}

// Info msg to std out and send an info event to the Sentry server
// pointed by `SENTRY_DSN` environment variable.
// /!\ This is a non blocking call, if the program die just after the call, the message can be lost
func Info(msg string) {
	InfoWithTags(msg, map[string]string{})
}

func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Error msg to std out and send an error event to the Sentry server
// pointed by `SENTRY_DSN` environment variable.
// tags can contain a map of additional tags information
// /!\ This is a non blocking call, if the program die just after the call, the message can be lost
func ErrorWithTags(err error, tags map[string]string) {
	raven.CaptureError(err, tags)
	log.Printf("%s, %s", "[ERROR]", err)
}

// Error msg to std out and send an error event to the Sentry server
// pointed by `SENTRY_DSN` environment variable.
// /!\ This is a non blocking call, if the program die just after the call, the message can be lost
func Error(err error) {
	ErrorWithTags(err, map[string]string{})
}

func FatalWithTags(msg string, tags map[string]string) {
	packet := &raven.Packet{
		Message: msg,
		Level:   raven.FATAL,
	}
	eventID, ch := raven.Capture(packet, tags)
	if eventID != "" {
		<-ch
	}
	log.Fatalf("%s, %s", "[FATAL]", msg)
}

func Fatal(msg string) {
	FatalWithTags(msg, map[string]string{})
}

func Fatalf(format string, a ...interface{}) {
	Fatal(fmt.Sprintf(format, a...))
}

func PanicWithTags(f func(), tags map[string]string) {
	raven.CapturePanicAndWait(f, tags)
}

func Panic(f func()) {
	PanicWithTags(f, map[string]string{})
}

// PrettyPrint convert any struct into formatted JSON like string.
// This is useful for debugging.
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
