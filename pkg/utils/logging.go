package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(file string) {
	logfile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("logfile error: %s", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiWriter)
}
