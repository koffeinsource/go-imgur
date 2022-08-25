package klogger

import (
	"context"

	"google.golang.org/appengine/log"
)

// GAELogger is a simple log wrapper for the GAE log
type GAELogger struct {
	Context context.Context
}

// Criticalf is like Debugf, but at Critical level and panics after printing the error.
func (l GAELogger) Criticalf(format string, args ...interface{}) {
	log.Criticalf(l.Context, format, args...)
}

// Debugf formats its arguments according to the format, analogous to fmt.Printf, and prints the text using log.Printf.
func (l GAELogger) Debugf(format string, args ...interface{}) {
	log.Debugf(l.Context, format, args...)
}

// Errorf is like Debugf, but at Error level.
func (l GAELogger) Errorf(format string, args ...interface{}) {
	log.Errorf(l.Context, format, args...)
}

// Infof is like Debugf, but at Info level.
func (l GAELogger) Infof(format string, args ...interface{}) {
	log.Infof(l.Context, format, args...)
}

// Warningf is like Debugf, but at Warning level.
func (l GAELogger) Warningf(format string, args ...interface{}) {
	log.Warningf(l.Context, format, args...)
}
