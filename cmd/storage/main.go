package main

import (
	"fmt"
	"log"

	"github.com/ShishVi/storage/internal/storage"
)

func main() {
	store := storage.NewStorage()

	file, err := store.Upload("test.txt", []byte("hello man!"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Upload file", file)
}
