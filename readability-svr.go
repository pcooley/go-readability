package main

import (
	"fmt"
	"net/http"
	"readability/readability"
)

type ReadabilityImpl struct{}

func (c ReadabilityImpl) Readability(paragraph string) {

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World, I love %s!", r.URL.Path[1:])
}

func main() {
	idl := barrister.MustParseIdlJson([]byte(readability.IdlJsonRaw))
	svr := readability.NewJSONServer(idl, true, ReadabilityImpl{})
	http.HandleFunc("/", &svr)
	fmt.Println("Starting Readability server on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
