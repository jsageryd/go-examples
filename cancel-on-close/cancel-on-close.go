package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", cancelOnClose(http.HandlerFunc(index)))
	addr := ":8080"
	fmt.Printf("Listening on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Waiting for possible cancellation...")
	select {
	case <-r.Context().Done():
		fmt.Println("Request cancelled")
	case <-time.After(5 * time.Second):
		fmt.Println("Request not cancelled")
	}
	io.WriteString(w, "Hello, World\n")
}

func cancelOnClose(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithCancel(r.Context())
			notify := w.(http.CloseNotifier).CloseNotify()
			go func() {
				<-notify
				cancel()
			}()
			h.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
