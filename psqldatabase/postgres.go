package psqldatabase

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //needed for postgres connector
	"github.com/nitsuan/cero_pwd_backend_go/data"
)

var host, user, password, dbname string
var port int64

//
func GetDatabaseEnv() {
	err := godotenv.Load("./environment/psqldatabase.env")
	if err != nil {
		panic(err.Error())
	}
	host, user, password, dbname = os.Getenv("HOST"), os.Getenv("DATABASE_USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME")
	port, _ = strconv.ParseInt(os.Getenv("PORT"), 10, 64)
}

func getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
}

//
func InsertIntoPwdColTable(name string, url string, username string, password string) {
	postgresConnector, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		panic(err.Error())
	}
	defer postgresConnector.Close()
	newUUID, err := uuid.NewRandom()
	if err != nil {
		panic(err.Error())
	}
	commandString := fmt.Sprintf("INSERT INTO password_collection VALUES ('%s', '%s', '%s', '%s', '%s')", newUUID.String(), name, url, username, password)
	_, err = postgresConnector.Exec(commandString)
	if err != nil {
		panic(err.Error())
	}
}

//
func ModifyDataPwdColTable(name, url, username, password, uuid string) {
	postgresConnector, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		panic(err.Error())
	}
	defer postgresConnector.Close()
	commandString := fmt.Sprintf("UPDATE password_collection SET name='%s', url='%s', username='%s', password='%s' WHERE uuid='%s'", name, url, username, password, uuid)
	_, err = postgresConnector.Exec(commandString)
	if err != nil {
		panic(err.Error())
	}
}

//
func DeletefromPwdColTable(uuid string) {
	postgressConnector, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		panic(err.Error())
	}
	defer postgressConnector.Close()
	commandString := fmt.Sprintf("DELETE FROM password_collection WHERE uuid='%s'", uuid)
	_, err = postgressConnector.Exec(commandString)
	if err != nil {
		panic(err.Error())
	}
}

//
func SelectfromPwdColTable() string {
	postgresConnector, err := sql.Open("postgres", getConnectionString())
	// pcr := make([]PasswordCollectionRow, 10)
	if err != nil {
		panic(err.Error())
	}
	defer postgresConnector.Close()
	fmt.Println("Connection Successful")
	rows, err := postgresConnector.Query("SELECT * FROM password_collection")
	if err != nil {
		panic(err.Error())
	}
	i := 0
	parsedString := "["
	for rows.Next() {
		tempPCR := data.PasswordCollectionRow{}
		var uuid, name, url, username, password string
		if err = rows.Scan(&uuid, &name, &url, &username, &password); err != nil {
			panic(err.Error())
		}
		tempPCR.SetValues(uuid, name, url, username, password)
		if i != 0 {
			parsedString += ","
		}
		parsedString += tempPCR.ToJSON()
		i++
	}
	parsedString += "]"
	return parsedString
}
