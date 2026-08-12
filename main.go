package main

import (
	"fmt"
	"os"
	"net/http"
)

func errh(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	filename := "./site" + r.URL.Path
	fmt.Printf("Sent: %s\n", r.URL.Path)
	
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		http.ServeFile(w, r, "./site/404.html")
	} else {
		http.ServeFile(w, r, filename)
	}
}

func main() {
	fmt.Println("// GoWebServer -- Made by SvZ_Code //")
	http.HandleFunc("/", handler)
	
	var port string
	
	if len(os.Args) == 1 {
		port = ":8089"
	} else {
		port = ":" + os.Args[1]
	}
	
	err := http.ListenAndServe(port, nil)
	errh(err)
}

