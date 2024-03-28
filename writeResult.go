package main

import (
	"log"
	"os"
)

func writeResult(result *[]WordWeight) {

	missingContainer := make(map[string]bool)

	analytics, out, missingLog, errorLog, closeFn := createResult()

	var err error

	defer func(out *os.File) {
		closeFn()
	}(analytics)

	//resultHeader := fmt.Sprintf("%"+strconv.Itoa(PrintWidthRowNum+1)+
	//	"s%"+strconv.Itoa(PrintWidthWord+1)+"s%s", "#,", "Név,", " Súlyok, Ismeretlen Betűk\n")
	resultHeader := "#,Név,Súlyok,Ismeretlen Betűk\n"
	_, err = analytics.WriteString(resultHeader)
	if err != nil {
		log.Fatal(err)
	}
	_, err = errorLog.WriteString(resultHeader)
	if err != nil {
		log.Fatal(err)
	}

	for _, weighted := range *result {
		_, err = analytics.WriteString(weighted.String(false) + "\n")
		if err != nil {
			log.Fatal(err)
		}
		_, err = out.WriteString(weighted.word.text + "\n")
		if err != nil {
			log.Fatal(err)
		}
		if len(weighted.missing) > 0 {
			_, err = errorLog.WriteString(weighted.String(true) + "\n")
			if err != nil {
				log.Fatal(err)
			}
			for i := range weighted.missing {
				missingContainer[weighted.missing[i]] = true
			}
		}
	}

	for missing := range missingContainer {
		_, err = missingLog.WriteString(missing + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
