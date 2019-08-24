package main

import "net/http"

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":80", nil)
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("success."))
}
