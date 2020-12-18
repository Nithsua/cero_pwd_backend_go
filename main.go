package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitsuan/cero_pwd_backend_go/environment"
	"github.com/nitsuan/cero_pwd_backend_go/psqldatabase"
)

func indexFunction(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	if parameters["action"][0] == "select" {
		fmt.Println("Action: Select")
	} else if parameters["action"][0] == "delete" {
		fmt.Println("Action: Delete")
		id, _ := strconv.ParseInt(parameters["id"][0], 10, 32)
		psqldatabase.DeletefromPwdColTable(int(id))
	} else if parameters["action"][0] == "modify" {
		fmt.Println("Action: Modify")

	} else if parameters["action"][0] == "create" {
		fmt.Println("Action: Create")
	}
	fmt.Fprintf(w, psqldatabase.SelectfromPwdColTable())
}

func main() {
	environment.LoadEnvironment()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/password_collection", indexFunction)
	log.Fatal(http.ListenAndServe(":8000", router))
}
