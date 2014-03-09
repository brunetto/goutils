package debug

import (
	"log"
	"time"
)

func timeMe(start time.Time) {
    elapsed := time.Since(start)
    log.Println("Wall time for ", FName(false), ": ", name, elapsed)
}