package logger

import (
	"log"
	"os"
)

var (
	Info = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	Err  = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
)
