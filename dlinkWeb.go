// dlinkWeb.go - dlink as a webservice
//
// 24.01.2014 - ICE 1090
//
// flw@posteo.de
//

package main

import(
    "log"
    "fmt"
	"regexp"
    "net/http"
    "io/ioutil"
    "unicode/utf8"
    "html/template"
    "net/http/httputil"
)

// constants and variables:
var templates = template.Must(
    template.ParseFiles(
        "tmpl/upload.html", 
        "tmpl/download.html",
    ),
)

type File struct {
    Name        string
    Size        int
    ContentType string
    Content     []byte
    Error       string
}

func upHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tmpl/upload.html")
    p := ""
    t.Execute(w, p)
}

func doHandler(w http.ResponseWriter, r *http.Request) {
    // get multipart-file from request
    mpr, mpHeader, err := r.FormFile("file")
    //mpr, _, err := r.FormFile("file")
    if err != nil {
        log.Fatal(err)
    }
    content, err := ioutil.ReadAll(mpr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    //Input Validation
    // - Input should not be over a specified size
    if len(content) >= 1024 * 1024 {
        http.Error(w, "Input to large", http.StatusInternalServerError)
        return
    }
    // - Input should just contain runes (valid Unicode Codepoints)
    if utf8.Valid(content) == false {
        //http.Error(w, "Input contains non unicode characters", http.StatusInternalServerError)
        t, _ := template.ParseFiles("tmpl/upload.html")
        p := File{}
        p.Error = "Input contains non unicode characters"
        t.Execute(w, p)
        return
    }
    // replace content
	re := regexp.MustCompile(`http(s)?:\/\/`) // matches 'http://' as well as 'https://'
	content = re.ReplaceAll(content, []byte("hXXp${1}://"))
    //fmt.Fprintf(w, "%s", content)
    //t, _ := template.ParseFiles("tmpl/download.html")
    t, _ := template.ParseFiles("tmpl/upload.html")
    p := File{}
    p.Content = content
    p.Size = len(content)
    p.ContentType = mpHeader.Header.Get("Content-Type")
    p.Name = mpHeader.Filename
    t.Execute(w, p)
}

func reqDumper(w http.ResponseWriter, r *http.Request) {
    dumpedReq, err := httputil.DumpRequest(r, true)
    if err != nil {
        log.Fatal("httputil.DumpRequestOut:", err)
    }
    log.Printf("%s\n", dumpedReq)
    fmt.Fprintf(w, "%s", dumpedReq)
}

func main() {
    http.HandleFunc("/", upHandler)
    http.HandleFunc("/up/", upHandler)
    http.HandleFunc("/do/", doHandler)
    http.HandleFunc("/dump/", reqDumper)
    http.ListenAndServe(":9090", nil)
}
