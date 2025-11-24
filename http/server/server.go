package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer() {

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		io.WriteString(w, "Finished!")
	}))
	mux.Handle("/sayHello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		word := request.URL.RawQuery
		io.WriteString(writer, fmt.Sprintf("you say: %s to me?", word))
	}))
	srv := &http.Server{Addr: ":8081", Handler: mux}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}
