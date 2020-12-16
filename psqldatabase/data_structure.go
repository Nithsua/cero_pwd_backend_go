package psqldatabase

import (
	"encoding/json"
	"fmt"
)

//
type PasswordCollectionRow struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//
func (pcr *PasswordCollectionRow) SetValues(id, name, url, username, password string) {
	pcr.ID, pcr.Name, pcr.URL, pcr.Username, pcr.Password = id, name, url, username, password
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
func (pcr *PasswordCollectionRow) FromJSON(jsonString string) {
	json.Unmarshal([]byte(jsonString), pcr)
}
