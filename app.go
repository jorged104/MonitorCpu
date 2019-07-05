package main

import (
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	w.Write([]byte(message))
}

func testw(w http.ResponseWriter, r *http.Request) {
	message := " Hola Mundo test "
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/test", testw)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
