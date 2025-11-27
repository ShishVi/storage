package main

import (
	"fmt"

	"github.com/ShishVi/storage/internal/storage"
)

func main() {
	store := storage.NewStorage()
	fmt.Println("This is storage prog", store)
}
