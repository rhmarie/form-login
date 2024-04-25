package main 

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := 8012
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Server running on Port %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}