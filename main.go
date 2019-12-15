package main

import "fmt"
import "net/http"

func main() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Dao")
	})

	http.ListenAndServe(":8080", nil)
}
