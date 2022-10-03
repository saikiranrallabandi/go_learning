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

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database= GetConnection()

	var error error
	var insert *sql.Stmt
	insert,error = database.Prepare("Insert INTO Customer(CustomerId,CustomerName,SSN) VALUES(?,?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerId, customer.CustomerName, customer.SSN)
	defer database.Close()

}

// Update method with parameter customer
func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE Customer SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	//log.Println("INSERT: Customer Name: " + customer.name + " | SSN: " + customer.ssn)

	defer database.Close()

	//return Customer{}
}

//delete method
func deleteCustomer(customer Customer){
	var database *sql.DB
	database= GetConnection()
	var error error
	var delete *sql.Stmt
	delete,error = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)
	defer database.Close()
}

func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println("Before Update Customers", customers)
	var customer Customer
	    customer.CustomerName = "saikiran"
	customer.SSN = "272432387"
	customer.CustomerId = 6


	//InsertCustomer(customer)
	//UpdateCustomer(customer)
	deleteCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Update", customers)
}