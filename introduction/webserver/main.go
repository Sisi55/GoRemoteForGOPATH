package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello world!")); err != nil {
			panic(err) // 이러면 되는건가 ?
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
