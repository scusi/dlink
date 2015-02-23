// dlink.go - opens any number of files and replaces 'http://' to 'hxxp://', in order to make links non clickable anymore.
// open any number of given files
// replace all links into non clickable versions within the content of each file
// write the file back to it's original filename
//
// flw@posteo.de
// 24.01.2014 - ICE 1090

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// checkErr a function to check the error return value of another function.
func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
}

// loader - returns the content of a given file as []byte
// takes a string (the filename of the file to read) as argument
// returns a []byte (containing content of file) and error
func loader(filename string) (content []byte, err error) {
	content, err = ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// writeFile takes an array of bytes, and a filename as a string as argument, writes
// those bytes into a file with the supplied name (filename), returns the number of bytes written
func writeFile(b []byte, filename string) (i int) {
	f, err := os.Create(filename)
	checkErr(err)
	defer f.Close()
	i, err = f.Write(b)
	checkErr(err)
	return i
}

// main takes arbitary filenames as arguments from command line and rewrites those files.
func main() {
	// compile the regex, make sure it is valid and does actually compile
	re := regexp.MustCompile(`hXXp(s)?:\/\/`) // matches 'http://' as well as 'https://'
	// get the arguments from the command line
	args := os.Args[1:]
	// iterate over arguments
	for i := 0; i < len(args); i++ {
		// get current filename from arguments
		filename := args[i]
		// load file content and check for errors
		content, err := loader(filename)
		checkErr(err)
		// replace what matches our regex
		//  replaces 'http://'  into 'hXXp://' or subsequent 'https://' into 'hXXps://'
		new_content := re.ReplaceAll(content, []byte("http${1}://"))
		// write content (with replacements) back to file
		i := writeFile(new_content, filename+".rlinked")
		// print out how many bytes we wrote onto what file
		fmt.Printf("%d bytes written to '%s'\n", i, filename+".rlinked")
	}
}
