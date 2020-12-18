package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitsuan/cero_pwd_backend_go/data"
	"github.com/nitsuan/cero_pwd_backend_go/environment"
	"github.com/nitsuan/cero_pwd_backend_go/psqldatabase"
)

func indexFunction(w http.ResponseWriter, r *http.Request) {
	parameters := r.URL.Query()
	if parameters["action"][0] == "select" {
		fmt.Println("Action: Select")
	} else if parameters["action"][0] == "delete" {
		fmt.Println("Action: Delete")
		uuid := parameters["uuid"][0]
		psqldatabase.DeletefromPwdColTable(uuid)
	} else if parameters["action"][0] == "modify" {
		fmt.Println("Action: Modify")
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			panic(err.Error())
		}
		tempPCR := data.PasswordCollectionRow{}
		tempPCR.FromJSON(body)
		psqldatabase.ModifyDataPwdColTable(tempPCR.Name, tempPCR.URL, tempPCR.Username, tempPCR.Password, tempPCR.UUID)
	} else if parameters["action"][0] == "create" {
		fmt.Println("Action: Create")
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			panic(err.Error())
		}
		tempPCR := data.PasswordCollectionRow{}
		tempPCR.FromJSON(body)
		psqldatabase.InsertIntoPwdColTable(tempPCR.Name, tempPCR.URL, tempPCR.Username, tempPCR.Password)
	}
	fmt.Fprintf(w, psqldatabase.SelectfromPwdColTable())
}

func main() {
	environment.LoadEnvironment()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/password_collection", indexFunction)
	log.Fatal(http.ListenAndServe(":8000", router))
}
