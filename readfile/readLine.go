package readfile

import (
	"bufio"

// 	"log"
)

func Readln(nReader *bufio.Reader) (string, error) {
	/*from http://stackoverflow.com/questions/6141604/go-readline-string*/

	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = nReader.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
