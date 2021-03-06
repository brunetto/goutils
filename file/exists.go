package file

// From http://stackoverflow.com/a/12527546
import (
	"os"
)

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
        }
    }
    return true
}
