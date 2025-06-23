package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	// Your code here
	mux := http.NewServeMux()

	svr := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		Handler: mux,
	}
	http2.ConfigureServer(svr, nil)

	actions(mux) // Register actions with the multiplexer

	fmt.Printf("Starting server on %s:%s\n", os.Getenv("HOST"), os.Getenv("PORT"))
	err := svr.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
