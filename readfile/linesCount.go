package readfile

import (
	"bufio"
	"log"
	"os"
)

func LinesCount(fileObj *os.File) (totCount int, voidCount int) {
	var (
		nReader  *bufio.Reader
		readLine string
		err      error
	)

	log.Println("Counting lines")

	nReader = bufio.NewReader(fileObj)

	totCount, voidCount = 0, 0
	for {
		if readLine, err = Readln(nReader); err != nil {
			break
		}
		if len(readLine) == 0 {
			voidCount++
		}
		totCount++
	}

	// 	log.Println("Rewind file")
	if _, err := fileObj.Seek(0, 0); err != nil {
		log.Fatal("Failed seek of file with error ", err)
	}
	// 	log.Println("Done")
	return totCount, voidCount
}
