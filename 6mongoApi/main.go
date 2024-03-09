package main

import (
	"fmt"
	"net/http"

	"github.com/TaniKroos/mongoApi/router"
)

func main() {
	r := router.Router()
	fmt.Println("Hello MOngo")
	fmt.Println("Server is getting started...")
	http.ListenAndServe(":4000", r)
	fmt.Println("listening at 4000")
}
