package database

import "github.com/aaabigfish/gopkg/log"

// Logger implements io.Writer interface to redirect SQLBoiler debug output to GCT logger
type Logger struct{}

// Write takes input and sends to GCT logger
func (l Logger) Write(p []byte) (n int, err error) {
	log.Debugf("Database SQL: %s", p)
	return 0, nil
}
