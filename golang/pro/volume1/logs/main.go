package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	basepath := "pro/volume1/logs"
	LOGFILE1 := filepath.Join(basepath, "mGo.log")

	f, err := os.OpenFile(LOGFILE1, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("can't open log file")
	}

	defer f.Close()

	iLog := log.New(f, "iLog ", log.LstdFlags|log.Llongfile)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 3rd edition!")
}
