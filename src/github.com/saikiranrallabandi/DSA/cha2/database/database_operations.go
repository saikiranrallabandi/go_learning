package main
import (
	"fmt"
	"log"
	"os"
	"database/util"
	"database/sql"
     _ "github.com/go-sql-driver/mysql"
)


// Customer Class
type Customer struct {
    CustomerId int
    CustomerName string
    SSN string
}

//GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB)  {
	config, err := util.LoadConfig(".")
	if err != nil {
        log.Fatal("cannot load config:", err)
    }
	fmt.Println(config.DBSource);
	database, error := sql.Open(config.DBDriver, os.Getenv("DB_SOURCE"))
	if error != nil {
		panic(error.Error())
	}
	return database

}

// Get Customers method returns Customer Array

func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}


	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
		panic(error.Error())
	}
	customer.CustomerId = customerId
	customer.CustomerName = customerName
	customer.SSN = ssn
	customers = append(customers, customer)
	}
     defer database.Close()

	 return customers
}

func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println("Customers", customers)
}