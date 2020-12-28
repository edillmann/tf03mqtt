package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type formatter struct {
}

func (f formatter) Format(e *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintln(e.Time.Format(time.StampMilli), " ", e.Message)), nil
}

func setLogLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.Warnln(err, "Fallback to info level")
		return
	}
	log.SetLevel(lvl)
}

