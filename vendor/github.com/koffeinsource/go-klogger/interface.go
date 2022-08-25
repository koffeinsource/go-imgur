package klogger

// KLogger is a minimalistic interface to wrap both the Google App Engine and default golang log packages
type KLogger interface {
	Criticalf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}
