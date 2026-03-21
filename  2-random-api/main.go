package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main(){
	router := http.NewServeMux()
	router.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		rn := rand.Intn(7)

		w.Write([]byte(fmt.Sprint(rn)))
		return
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}