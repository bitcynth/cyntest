package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

var buildVersion string
var buildTime string

type indexPageData struct {
	Hostname     string
	BuildTime    string
	BuildVersion string
}

func main() {
	r := mux.NewRouter()
	hostname, _ := os.Hostname()

	indexTmpl := template.Must(template.ParseFiles("templates/index.tmpl.html"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := indexPageData{
			Hostname:     hostname,
			BuildTime:    buildTime,
			BuildVersion: buildVersion,
		}
		indexTmpl.Execute(w, data)
	})

	r.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "version: %s\nbuild_time: %s", buildVersion, buildTime)
	})

	http.ListenAndServe(":5000", r)
}
