package file

// https://groups.google.com/d/msg/golang-nuts/JNyQxQLyf5o/kbGnTUK32TkJ
import ( 
	"log"
	"io" 
	"os" 
) 

func CopyFile(src, dst string) (int64, error) { 
	var(
		sf *os.File
		df *os.File
		err error
	)
	if sf, err = os.Open(src); err != nil {/*return 0,*/ log.Fatal("On opening source: ", err)} 
	defer sf.Close() 
	if df, err = os.Create(dst); err != nil {/*return 0,*/ log.Fatal("On opening destination: ", err)} 
	defer df.Close() 
	return io.Copy(df, sf) 
} 

/* SMART EXAMPLE OF USE
func main() { 
        fn := "copyfile.go" 
        n, err := CopyFile("(copy of) "+fn, fn) 
        if err != nil { 
                fmt.Println(n, err) 
        } 
} */