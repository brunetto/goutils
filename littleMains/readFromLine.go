package main

// Example provided by Egon here:
// https://groups.google.com/forum/#!topic/golang-nuts/-a4g5Jde8Ik%5B1-25-false%5D
// http://play.golang.org/p/72CWT6-kBU

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
)

type Line struct {
	Fields []string
	Error  error
}

func (line *Line) handle(err error) {
	if err != nil && line.Error == nil {
		line.Error = err
	}
}

func (line *Line) Float64(i int) float64 {
	r, err := strconv.ParseFloat(line.Fields[i], 64)
	line.handle(err)
	return r
}

func (line *Line) String(i int) string {
	return line.Fields[i]
}

func (line *Line) Int(i int) (r int) {
	r, err := strconv.Atoi(line.Fields[i])
	line.handle(err)
	return r
}

type Info struct {
	A float64
	B string
	C float64
	D int
}

type Infos []Info

func (info *Info) ReadFrom(line *Line) {
	info.A = line.Float64(0)
	info.B = line.String(1)
	info.C = line.Float64(2)
	info.D = line.Int(3)
}

const example = `124.21,hello,12.41,521
124.21,hello,12.41,521
hello,hello,12.41,521
124.21,hello,12.41,521
`

func main() {
	reader := csv.NewReader(bytes.NewBufferString(example))
	reader.Comma = ',' // alternatively use '\t'
	rows, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	infos := make([]Info, len(rows))
	for i, row := range rows {
		line := &Line{row, nil}
		(&infos[i]).ReadFrom(line)
		if line.Error != nil {
			fmt.Printf("line %v: parse error: %v\n", i, line.Error)
		}
	}
}
