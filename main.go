package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	ctx := context.Background()
	root := weaver.Init(ctx)

	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hello listener available on %v\n", lis)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Hello, %s!\n", reversed)
	})
	http.Serve(lis, nil)
}
