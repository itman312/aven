package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>ni ..hhhhao ma ...80zhende ?</h1>")
}

func main(){
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":80", nil)
}