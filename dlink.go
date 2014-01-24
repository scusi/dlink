// go prog to replace 'http://' into 'hXXp://' for non-clickable links
//
// flw@posteo.de
//
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // we look for 'http://' or 'https://' with the following regex
    re := regexp.MustCompile(`http(s)?:\/\/`)

    // print tests for our above regex
    fmt.Println("http://heise.de       : MatchString : ", re.MatchString("Der link ist: http://www.heise.de."))
    fmt.Println("https://scusiblog.org : MatchString : ", re.MatchString("Kann es auch: https://scusiblog.org?"))

    // let's replace the found strings partly
    fmt.Println(re.ReplaceAllString("Malware Site: http://eorohfoi.ldshfodsf.com/6543234567?98765=9876543", "hXXp://") )
    fmt.Println(re.ReplaceAllString("Malware Site: https://eorohfoi.ldshfodsf.com/6543234567?98765=9876543", "hXXps://") )
}
