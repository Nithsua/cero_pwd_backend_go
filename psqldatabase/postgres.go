package psqldatabase

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //needed for postgres connector
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
	commandString := fmt.Sprintf("INSERT INTO password_collection(name, url, username, password) VALUES ('%s', '%s', '%s', '%s')", name, url, username, password)
	_, err = postgresConnector.Exec(commandString)
	if err != nil {
		panic(err.Error())
	}
}

//
func ModifyDataPwdColTable(name string, url string, username string, password string, id int) {
	postgresConnector, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		panic(err.Error())
	}
	defer postgresConnector.Close()
	commandString := fmt.Sprintf("UPDATE table password_collection SET name='%s', url='%s', username='%s', password='%s WHERE uid='%d'", name, url, username, password, id)
	_, err = postgresConnector.Exec(commandString)
	if err != nil {
		panic(err.Error())
	}
}

//
func DeletefromPwdColTable(id int) {
	postgressConnector, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		panic(err.Error())
	}
	defer postgressConnector.Close()
	commandString := fmt.Sprintf("DELETE FROM password_collection WHERE uid=%d", id)
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
		tempPCR := PasswordCollectionRow{}
		var id, name, url, username, password string
		if err = rows.Scan(&id, &name, &url, &username, &password); err != nil {
			panic(err.Error())
		}
		tempPCR.SetValues(id, name, url, username, password)
		if i != 0 {
			parsedString += ","
		}
		parsedString += tempPCR.ToJSON()
		i++
	}
	parsedString += "]"
	return parsedString
}
