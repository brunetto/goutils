package main

import (
	"bufio"
		"fmt"
	"log"
	"os"
	"path/filepath"
// 	"reflect"
	"regexp"
	"sort"
	"strconv"
	// 	"strings"
	"time")

func main() {
	
	tGlob0 := time.Now()
	
	// Declare variables
	var inPath string
	var inFile []string
	var written ResultsMap
	var oral ResultsMap
	var final ResultsMap
	
	inPath = "./"
	inFile = []string{"scritto.txt", "orale.txt"}
	
	
	written = ReadWritten(inPath, inFile[0])
	
	oral = ReadOral(inPath, inFile[1])
	
	// Print results
	/*
	fmt.Println("###################################")
	log.Println("Written results:")
	fmt.Println("###################################")
	written.PrintMap()
	
	fmt.Println("###################################")
	log.Println("Written results:")
	fmt.Println("###################################")
	oral.PrintMap()
	*/
	
	final = make(ResultsMap)
	
	keys := written.Keys()
	for _, key := range keys {
		if _, exists := oral[key]; !exists {
			continue
		}
		final[key] = written[key] + oral[key]
	}
	//Print final results sorted by key
// 	final.PrintMap()
	
// 	fmt.Println(final.Values())
// 	final.PrintMapVal()

	sortedResults :=sortMapByValueDescending(final)
	
	for idx, value := range(sortedResults) {
		fmt.Println(idx+1, value.Key, value.Value)
	}
	
	tGlob1 := time.Now()
	log.Println("\nWall time for all ", tGlob1.Sub(tGlob0))
}



func ReadWritten(inPath string, inFile string) (written ResultsMap) {
	// Function variables
	var fileObj *os.File
	var nReader *bufio.Reader
	var readLine string
	var err error
	var writtenRegexp *regexp.Regexp
	var regexResult []string
	var name string
	var points int64
	
	// Init regexp
	writtenRegexp = regexp.MustCompile(`\d{1,2}\s(\D+)\s\d+/\d+/\d+\s(\d+)`)
	
	// Open the file
	if fileObj, err = os.Open(filepath.Join(inPath, inFile)); err != nil {
		log.Fatal("Can't open: error: ", err)
	}
	defer fileObj.Close() // To be done before exit
	
	// Create a reader to read the file
	nReader = bufio.NewReader(fileObj)
	
	// Init map
	written = make(ResultsMap) 
	
	nLine := 1
	for {
		if readLine, err = nReader.ReadString('\n'); err != nil {
			log.Println("Done reading ", nLine, " lines from file with err", err)
			break
		}
// 		fmt.Printf("Read %v on line %v\n", readLine[:len(readLine)-1], nLine)
		if regexResult = writtenRegexp.FindStringSubmatch(readLine); regexResult == nil {
			log.Println("With regexp ", writtenRegexp)
			log.Fatal("no match, nil regex result on line ", nLine)
		}
		// Regexp results
		name = regexResult[1]
		points, _ = strconv.ParseInt(regexResult[2], 10, 64)
		
		// Store results
		written[name] = points
		
		// Increment line
		nLine++
	}
	return written
}

func ReadOral(inPath string, inFile string) (oral ResultsMap) {
	// Function variables
	var fileObj *os.File
	var nReader *bufio.Reader
	var readLine string
	var err error
	var writtenRegexp *regexp.Regexp
	var regexResult []string
	var name string
	var points int64
	
	// Init regexp
	writtenRegexp = regexp.MustCompile(`(\D+)\s(\d+)`)
	
	// Open the file
	if fileObj, err = os.Open(filepath.Join(inPath, inFile)); err != nil {
		log.Fatal("Can't open: error: ", err)
	}
	defer fileObj.Close() // To be done before exit
	
	// Create a reader to read the file
	nReader = bufio.NewReader(fileObj)
	
	// Init map
	oral = make(ResultsMap) 
	
	nLine := 1
	for {
		if readLine, err = nReader.ReadString('\n'); err != nil {
			log.Println("Done reading ", nLine, " lines from file with err", err)
			break
		}
// 		fmt.Printf("Read %v on line %v\n", readLine[:len(readLine)-1], nLine)
		if regexResult = writtenRegexp.FindStringSubmatch(readLine); regexResult == nil {
			log.Println("With regexp ", writtenRegexp)
			log.Fatal("no match, nil regex result on line ", nLine)
		}
		// Regexp results
		name = regexResult[1]
		points, _ = strconv.ParseInt(regexResult[2], 10, 64)
		
		// Store results
		oral[name] = points
		
		// Increment line
		nLine++
	}
	return oral
}

type ResultsMap map[string]int64

// Retrieve and sort map keys
func (resultMap ResultsMap) Keys() (keys []string) {
	keys = make([]string, len(resultMap))
	idx := 0 
	for key, _ := range resultMap {
        keys[idx] = key
        idx++
    }
    sort.Strings(keys)
	return keys
}

func (resultMap ResultsMap) PrintMap() () {
	// Retrieve map sorted keys
	keys := resultMap.Keys()
	
	for _, key := range keys {
		value := resultMap[key] 
		fmt.Println(key, value)
	}
}

func (resultMap ResultsMap) Values() (values []int64) {
	values = make([]int64, len(resultMap))
	
	idx := 0
	for _, value := range resultMap {
		values[idx] = value
		idx++
	}
	sort.Sort(int64arr(values))
	return values
}

// func (resultMap ResultsMap) PrintMapVal() () {
// 	// Retrieve map sorted values
// 	values := resultMap.Values()
// 	
// 	
// 	
// 	
// 	idx := 0
// 	for _, value := range values {
// 		key := tempMap[value] 
// 		fmt.Println(idx, key, value)
// 		idx++
// 	}	
// }


type int64arr []int64
func (a int64arr) Len() int { return len(a) }
func (a int64arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }


// A data structure to hold a key/value pair.
type Pair struct {
  Key string
  Value int64
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairListAscend []Pair
func (p PairListAscend) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairListAscend) Len() int { return len(p) }
func (p PairListAscend) Less(i, j int) bool { return p[i].Value < p[j].Value }

type PairListDescend []Pair
func (p PairListDescend) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairListDescend) Len() int { return len(p) }
func (p PairListDescend) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it. 
func sortMapByValueAscending(m map[string]int64) PairListAscend {
   p := make(PairListAscend, len(m))
   i := 0
   for k, v := range m {
      p[i] = Pair{k, v}
      i++
   }
   sort.Sort(p)
   return p
}

// A function to turn a map into a PairList, then sort and return it. 
func sortMapByValueDescending(m map[string]int64) PairListDescend {
   p := make(PairListDescend, len(m))
   i := 0
   for k, v := range m {
      p[i] = Pair{k, v}
      i++
   }
   sort.Sort(p)
   return p
}







