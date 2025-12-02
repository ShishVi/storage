package main

import (
	"fmt"
	"log"

	"github.com/ShishVi/storage/internal/storage"
)

func main() {
	store := storage.NewStorage()

	file1, err1 := store.Upload("test.txt", []byte("hello man!"))
	if err1 != nil {
		log.Fatal(err1)
	}
	file2, err2 := store.Upload("test2.txt", []byte("hello man!!!"))
	if err2 != nil {
		log.Fatal(err1)
	}
	file3, err3 := store.Upload("test3.txt", []byte("hello woman!"))
	if err3 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Upload file", file1)
	fmt.Println("Upload file", file2)
	fmt.Println("Upload file", file3)

	foundFile, ok := store.GetByID(file3.ID)
	if ok != nil {
		log.Fatal(ok)
	}
	fmt.Println("Found file", foundFile)
}
