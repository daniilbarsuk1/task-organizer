package test

import (
	"net/http"
)

func test() {
	err := http.ListenAndServe("localhost:8080", http.NewServeMux())
	if err != nil {
		panic(err)
	}
}
