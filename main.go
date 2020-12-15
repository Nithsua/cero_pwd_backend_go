package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitsuan/cero_pwd_backend_go/environment"
	"github.com/nitsuan/cero_pwd_backend_go/psqldatabase"
)

func indexFunction(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	fmt.Println(parameters)
	fmt.Fprintf(w, "Test, its listen and server from go")
}

func main() {
	environment.LoadEnvironment()
	psqldatabase.DatabaseConnectionTest()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/password_collection", indexFunction)
	log.Fatal(http.ListenAndServe(":8000", router))
}
