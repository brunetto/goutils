package debug

import (
    "log"
)

func LogDebug(debug bool, args ...interface{}) () {
	if debug {
		log.Println(args...)
	}
}
