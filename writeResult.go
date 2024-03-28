package main

import (
	"log"
	"os"
)

func writeResult(result *[]WordWeight) {

	missingContainer := make(map[string]bool)

	out, pure, missingLog, errorLog, closeFn := createResult()

	var err error

	defer func(out *os.File) {
		closeFn()
	}(out)

	//resultHeader := fmt.Sprintf("%"+strconv.Itoa(PrintWidthRowNum+1)+
	//	"s%"+strconv.Itoa(PrintWidthWord+1)+"s%s", "#,", "Név,", " Súlyok, Ismeretlen Betűk\n")
	resultHeader := "#,Név,Súlyok,Ismeretlen Betűk\n"
	_, err = out.WriteString(resultHeader)
	if err != nil {
		log.Fatal(err)
	}
	_, err = errorLog.WriteString(resultHeader)
	if err != nil {
		log.Fatal(err)
	}

	for _, weighted := range *result {
		_, err = out.WriteString(weighted.String(false) + "\n")
		if err != nil {
			log.Fatal(err)
		}
		_, err = pure.WriteString(weighted.word.text + "\n")
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
