package main

import (
	"log"
	"os"
)

func createResult() (*os.File, *os.File, *os.File, *os.File, func()) {
	// create analytics output file
	analytics, err := os.OpenFile(os.Args[3]+"_analytics.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// create out output file
	out, err := os.OpenFile(os.Args[3]+".txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// create file for missing characters
	missing, errUnknown := os.OpenFile(os.Args[3]+"_unknown.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if errUnknown != nil {
		log.Fatal(errUnknown)
	}

	// create file for errors
	errorLog, errLog := os.OpenFile(os.Args[3]+"_error.csv", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if errLog != nil {
		log.Fatal(errLog)
	}

	closeFunc := func() {
		err := analytics.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = out.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = missing.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = errorLog.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	return analytics, out, missing, errorLog, closeFunc
}
