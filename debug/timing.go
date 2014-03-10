package debug

import (
	"log"
	"time"
)

func TimeMe(start time.Time) {
    elapsed := time.Since(start)
    log.Println("Wall time for ", FName(false), ": ", elapsed)
}