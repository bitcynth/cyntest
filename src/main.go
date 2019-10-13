package main

import (
	"fmt"
	"net/http"
)

var buildVersion string
var buildTime string

func main() {
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "version: %s\nbuild_time: %s", buildVersion, buildTime)
	})
	http.ListenAndServe(":5000", nil)
}
