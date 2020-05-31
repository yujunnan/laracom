package api

import (
	"log"
	"net/http"
)

//noinspection ALL
func StartWebServer(port string) {
	log.Println("Starting HTTP service at port " + port)
	router := NewRouter()
	http.Handle("/", router)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}