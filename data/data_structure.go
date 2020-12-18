package data

import (
	"encoding/json"
	"fmt"
)

//
type PasswordCollectionRow struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//
func (pcr *PasswordCollectionRow) SetValues(uuid, name, url, username, password string) {
	pcr.UUID, pcr.Name, pcr.URL, pcr.Username, pcr.Password = uuid, name, url, username, password
}

//
func (pcr PasswordCollectionRow) ToJSON() string {
	data, err := json.Marshal(pcr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))
	return string(data)
}

//
func (pcr *PasswordCollectionRow) FromJSON(jsonData []byte) {
	json.Unmarshal(jsonData, pcr)
}
