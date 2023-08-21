package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world!")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(fmt.Sprintf("Bytes Written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
