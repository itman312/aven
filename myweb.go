package main

import (
	"io"
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>ni hao ma 80 ?</h1>")
}

func main(){
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":80", nil)
}