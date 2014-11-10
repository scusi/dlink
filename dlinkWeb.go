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
        //"tmpl/download.html",
    ),
)

type File struct {
    Name				string
    Size				int
    ContentType			string
    Content				[]byte
    ModifiedContent     []byte
    Error				string
}

func upHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tmpl/upload.html")
    p := ""
    t.Execute(w, p)
}

func readMultipartMimeFile(r *http.Request)(File){
    // get multipart-file from request
    mpr, mpHeader, err := r.FormFile("file")
    //mpr, _, err := r.FormFile("file")
    if err != nil {
        log.Fatal(err)
    }
    content, err := ioutil.ReadAll(mpr)
    if err != nil {
        //http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
    }
    p := File{}
    p.Content = content
    p.Size = len(content)
    p.ContentType = mpHeader.Header.Get("Content-Type")
    p.Name = mpHeader.Filename
	return p
}

func ValidateContent(w http.ResponseWriter, p File){
    //Input Validation
    // - Input should not be over a specified size
    if len(p.Content) >= 1024 * 1024 {
		log.Printf("Input Length of '%d' bytes is to large\n", len(p.Content))
        http.Error(w, "Input to large", http.StatusInternalServerError)
        return
    }
    // - Input should just contain runes (valid Unicode Codepoints)
    if utf8.Valid(p.Content) == false {
        //http.Error(w, "Input contains non unicode characters", http.StatusInternalServerError)
		log.Printf("Input contains non unicode characters\n")
        t, _ := template.ParseFiles("tmpl/upload.html")
        p := File{}
        p.Error = "Input contains non unicode characters"
        t.Execute(w, p)
        return
    }
}

func doHandler(w http.ResponseWriter, r *http.Request) {
	p := readMultipartMimeFile(r)
    //Input Validation
	ValidateContent(w, p)
    // replace content
	re := regexp.MustCompile(`http(s)?:\/\/`) // matches 'http://' as well as 'https://'
	p.ModifiedContent = re.ReplaceAll(p.Content, []byte("hXXp${1}://"))
    //t, _ := template.ParseFiles("tmpl/download.html")
    t, _ := template.ParseFiles("tmpl/upload.html")
    t.Execute(w, p)
}

func undoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("entering undoHandler...")
	log.Printf("Request: %#v\n", r)
	content:= []byte(r.FormValue("file"))
    p := File{ "file", len(content), "text/plain", content, []byte{}, ""}
	log.Printf("FormValue 'file': %s\n", p.Content)
    //Input Validation
	ValidateContent(w, p)
    // replace content
	log.Printf("perapring replacements, ")
	re := regexp.MustCompile(`http(s)?:\/\/`) // matches 'http://' as well as 'https://'
	log.Printf("replace\n")
	p.ModifiedContent = re.ReplaceAll(p.Content, []byte("hXXp${1}://"))
	p.Error = ""
    //fmt.Fprintf(w, "%s", content)
    //t, _ := template.ParseFiles("tmpl/download.html")
	log.Printf("preparing template\n")
    t, _ := template.ParseFiles("tmpl/upload.html")
	log.Printf("executing template content size: '%d'\n", p.Size)
    t.Execute(w, p)
}

func reHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("entering reHandler...")
	log.Printf("Request: %#v\n", r)
	content:= []byte(r.FormValue("file"))
    p := File{ "file", len(content), "text/plain", content, []byte{}, ""}
	log.Printf("FormValue 'file': %s\n", p.Content)
    //Input Validation
	ValidateContent(w, p)
    // replace content
	log.Printf("perapring replacements, ")
	re := regexp.MustCompile(`hXXp(s)?:\/\/`) // matches 'http://' as well as 'https://'
	log.Printf("replace\n")
	p.ModifiedContent = re.ReplaceAll(p.Content, []byte("http${1}://"))
	p.Error = "rlink"
    //fmt.Fprintf(w, "%s", content)
    //t, _ := template.ParseFiles("tmpl/download.html")
	log.Printf("preparing template\n")
    t, _ := template.ParseFiles("tmpl/upload.html")
	log.Printf("executing template content size: '%d'\n", p.Size)
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
    http.HandleFunc("/undo/", undoHandler)
    http.HandleFunc("/re/", reHandler)
    http.HandleFunc("/dump/", reqDumper)
    http.ListenAndServe(":9090", nil)
}
