package readfile

import (
	"os"
	"testing"
)

func Test_LinesCount(t *testing.T) {
	var (
		fileObj   *os.File
		voidCount int = 0
		totCount  int = 0
		err       error
	)

	if fileObj, err = os.Open("test.dat"); err != nil {
		panic(err)
	}
	defer fileObj.Close()

	totCount, voidCount = LinesCount(fileObj)

	if totCount != 6 {
		t.Error("Wrong total number of lines: ", totCount, " instead of ", 6) // log error if it did not work as expected
	} else {
		t.Log("Test passed.") // log some info if you want
	}
	if voidCount != 2 {
		t.Error("Wrong number of void lines: ", voidCount, " instead of ", 2) // log error if it did not work as expected
	} else {
		t.Log("Test passed.") // log some info if you want
	}
}
