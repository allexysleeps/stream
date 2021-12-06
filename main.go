package main

import (
	"flag"
	"github.com/allexysleeps/stream/ws"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveStreamerApp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/streamer" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "apps/streamer.html")
}

func serveWatcherApp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "apps/main.html")
}

func main() {
	hub := ws.NewHub()

	http.HandleFunc("/streamer", serveStreamerApp)
	http.HandleFunc("/", serveWatcherApp)

	http.HandleFunc("/stream", ws.ServeWs(hub))

	log.Printf("Server is running on %v\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
