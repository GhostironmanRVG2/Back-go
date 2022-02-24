package main

import (
	"net/http"
	"package/Routes"
)

func main() {
	//LOAD ROUTES
	http.Handle("/imagesget/", http.StripPrefix("/imagesget/", http.FileServer(http.Dir("./temp"))))
	Routes.ImagemRoute()
	Routes.LoadTests()
	//START THE SERVER
	err := http.ListenAndServe(":8080", nil)
	//error handler
	if err != nil {
		panic(err)
	}

}
