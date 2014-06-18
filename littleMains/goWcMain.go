package main

import (
	"fmt"
	"bitbucket.org/brunetto/goutils/readfile"
	"os"
)

func main() {
	
	var ( 
		fileObj *os.File
		voidCount int = 0
		totCount int = 0
		err error
	)
	
	if fileObj, err = os.Open(os.Args[1]); err != nil {
		panic(err)
	}
	defer fileObj.Close()
/*
	nReader = bufio.NewReader(fileObj)
	
	for {
		if readLine, err = readfile.Readln(nReader); err != nil {break}
		if len(readLine) == 0 {voidCount++}
		totCount++
	}
*/
	totCount, voidCount = readfile.LinesCount(fileObj)
	fmt.Println("Lines: ", totCount, " of which void: ", voidCount)
}
