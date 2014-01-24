// open a given file 
// replace all links into non clickable versions within the content of the file
// write the file back to it's original filename
//
// flw@posteo.de
// 24.01.2014 - ICE 1090

package main

import(
    "os"
    "io/ioutil"
    "fmt"
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
// returns a []byte (content of file) and error
func loader(filename string) (content []byte, err error) {
    content, err = ioutil.ReadFile(filename)
    if err != nil {
         return nil, err
    }
    return content, nil
}

// writeFile takes an array of bytes, and a filename as argument, writes 
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
    re := regexp.MustCompile(`http(s)?:\/\/`)
    args := os.Args[1:]
    for i := 0; i<len(args); i++ {
        filename := args[i] 
        content, err := loader(filename)
        checkErr(err)
        new_content := re.ReplaceAll(content, []byte("hXXp${1}://"))
	i := writeFile(new_content, filename)
	fmt.Printf("%d bytes written to '%s'\n", i, filename)
    }
}

