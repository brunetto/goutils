package readfile

import (
	"bufio"

// 	"log"
	"log"
	"strings"
	"strconv"
)

func Readln(nReader *bufio.Reader) (string, error) {
	/*from http://stackoverflow.com/questions/6141604/go-readline-string*/

	var (
		isPrefix bool  = true
		err      error = nil
		row, ln []byte
	)

	for isPrefix && err == nil {
		row, isPrefix, err = nReader.ReadLine()
		ln = append(ln, row...)
	}
	return string(ln), err
}

type LineReader struct {
	Reader *bufio.Reader
}

func NewLineReader  (r *bufio.Reader) LineReader {
	return LineReader{Reader:r}
}

func (l *LineReader) ReadDieOnErr () (string, bool) {
	row, err := Readln(l.Reader)
	if err != nil {
		if err.Error() != "EOF" {
			log.Fatal(err)
		}
		return row, true
	}
	return row, false
}

func (l *LineReader) ReadRetErr () (string, error, bool){
	var (
		row string
		err error
	)
	row, err = Readln(l.Reader)
	switch  {
	case err == nil:
		return row, err, false
	case err.Error() == "EOF":
		return row, err, true
	default:
		return row, err, false

	}
}

func (l *LineReader) LineToIntSlice (sep string) ([]int, error, bool){
	var (
		row string
		dataSlice = []int{}
		err error
		eof bool
		item string
		num int
	)
	row, err, eof = l.ReadRetErr()
	if err != nil {
		return []int{}, err, eof
	}

	for _, item = range strings.Split(row, sep) {
		num, err = strconv.Atoi(item)
		if err != nil {
			return []int{}, err, false
		}
		dataSlice = append(dataSlice, num)
	}

	return dataSlice, err, eof
}

func (l *LineReader) LineToInt64Slice (sep string) ([]int64, error, bool){
	var (
		row string
		dataSlice = []int64{}
		err error
		eof bool
		item string
		num int64
	)
	row, err, eof = l.ReadRetErr()
	if err != nil {
		return []int64{}, err, eof
	}

	for _, item = range strings.Split(row, sep) {
		num, err = strconv.ParseInt(item, 10, 64)
		if err != nil {
			return []int64{}, err, false
		}
		dataSlice = append(dataSlice, num)
	}

	return dataSlice, err, eof
}

func (l *LineReader) LineToFloat64Slice (sep string) ([]float64, error, bool){
	var (
		row string
		dataSlice = []float64{}
		err error
		eof bool
		item string
		num float64
	)
	row, err, eof = l.ReadRetErr()
	if err != nil {
		return []float64{}, err, eof
	}

	for _, item = range strings.Split(row, sep) {
		num, err = strconv.ParseFloat(item,64)
		if err != nil {
			return []float64{}, err, false
		}
		dataSlice = append(dataSlice, num)
	}

	return dataSlice, err, eof
}



