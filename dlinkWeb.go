// dlink as a webservice
//
// 24.01.2014 - ICE 1090
//
// flw@posteo.de
//

package main

import(
    "io/ioutil"
    "net/http"
    "html/template"
    "regexp"
    "errors"
)

// constants and variables:
var templates = template.Must(template.ParseFiles("tmpl/upload.html", "tmpl/download.html"))

func upHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tmpl/upload.html")
    p :=  ""
    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/up/", upHandler)
    //http.HandleFunc("/do/", doHandler)
    http.ListenAndServer(":9090", nil)
}
