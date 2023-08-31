package main

import (
	"errors"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/div", Divide)
	fmt.Printf("Starting app : %s%s", "http://localhost", port)
	_ = http.ListenAndServe(port, nil)
}

//
//
//
//
//
//
//
//

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := DeivdeValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "ERROR!!!")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f div by = %f is %f", 100.0, 10.0, f))
}

// addValue 앞문자가 소문자이면 패키지 프라이빗임
func addValue(x, y int) (int, error) {
	return x + y, nil
}

func DeivdeValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot div by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
