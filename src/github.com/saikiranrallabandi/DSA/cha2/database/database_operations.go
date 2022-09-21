package main
import (
	"fmt"
	"database/sql"
     _ "github.com/go-sql-driver/mysql"
	 "os"
    "strings"
)


// Customer Class
type Customer struct {
    CustomerId int
    CustomerName string
    SSN string
}

//GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB)  {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "mauFJcuf5dhRMQrjj"
	databaseName := "quotes"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@tcp(127.0.0.1:3306)/"+databaseName)
	if error != nil {
		panic(error.Error())
	}

	defer database.Close()

	return database

}

func main() {
	fmt.Println(GetConnection());
}