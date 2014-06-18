package main

import (
    "os"
    //"io/ioutil"
    "log"
    "fmt"
    "time"
    "os/exec"
    "regexp"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	// set ip download command
	command := "/usr/bin/curl"
	args := "http://checkip.dyndns.org"
	
	// compile regexp
	var digitsRegexp = regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	
	// loop to update the ip, infinite loop
    for ;;{
		// download ip string
		out, err := exec.Command(command, args).Output()
		if err != nil {
			log.Fatal(err)
		}
		
		// out is a binary buffer, convert into string
		ipString := string(out)
		
		// search for ip
		ipRes := digitsRegexp.FindString(ipString)
		fmt.Println(ipRes)
		
		// convert string to byte (found a better method)
		//ipByte := []byte(ipRes)
		//ioutil.WriteFile("ip.dat", ipByte, 0644)
		
		// create a file, it implements the Writer interface
		f, err := os.Create("ip.dat")
		
		// check for errors
		check(err)
		
		// close file before exit in case of problems
		defer f.Close()
		
		// write the string, discard (_) the number of bytes written
		_, err = f.WriteString(ipRes)
		// flush 
		f.Sync()
		f.Close()
		// wait
		time.Sleep(5 * time.Second)
    }
}
