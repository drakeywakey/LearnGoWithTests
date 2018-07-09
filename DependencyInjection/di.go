package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//Greet(os.Stdout, "DrakeyWakey")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Sup, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
