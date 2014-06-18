package debug

// https://groups.google.com/d/msg/golang-nuts/N2szzAFrcFM/s2WosyIDNP4J

import (
	"fmt"
	"log"
	"runtime"
)

// Whoami print the name of the calling function.
func FName(printName bool) string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}
	if printName {
		fmt.Println("###########################################")
		fmt.Print("###\t\t\t")
		log.Println(me.Name())
		fmt.Println("###########################################")
	}
	return me.Name()
}