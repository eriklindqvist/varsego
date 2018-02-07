package main

import (
	"log"
	"net/http"
	l "github.com/eriklindqvist/recepies_api/app/lib"
)

func main() {
	log.Printf("Server started")

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(l.Getenv("FILEBASE", "files")))))

	if err := http.ListenAndServe(":" + l.Getenv("PORT", "3003"), nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
