package server

import (
	"fmt"
	"net/http"
	"os"
)

// Index
func IndexHandle(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello friend")
}

func CallRoutes() {
	http.HandleFunc("/", IndexHandle)
}

func StartServer() {

	// import routes
	CallRoutes()
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
