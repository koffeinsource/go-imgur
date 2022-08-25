// +build !appengine

package klogger

import "log"

// CLILogger is a simple logger based on the the official go package
type CLILogger struct {
}

// Criticalf is like Debugf, but at Critical level and panics after printing the error.
func (l CLILogger) Criticalf(format string, args ...interface{}) {
	log.Printf("Critical: "+format, args...)
	panic("*** critical error ***")
}

// Debugf formats its arguments according to the format, analogous to fmt.Printf, and prints the text using log.Printf.
func (l CLILogger) Debugf(format string, args ...interface{}) {
	log.Printf("Debug: "+format, args...)
}

// Errorf is like Debugf, but at Error level.
func (l CLILogger) Errorf(format string, args ...interface{}) {
	log.Printf("Error: "+format, args...)
}

// Infof is like Debugf, but at Info level.
func (l CLILogger) Infof(format string, args ...interface{}) {
	log.Printf("Info: "+format, args...)
}

// Warningf is like Debugf, but at Warning level.
func (l CLILogger) Warningf(format string, args ...interface{}) {
	log.Printf("Warning: "+format, args...)
}
