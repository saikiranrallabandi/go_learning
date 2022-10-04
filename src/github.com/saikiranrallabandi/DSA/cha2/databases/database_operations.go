package main
import (
	"fmt"
	"log"
	"os"
	"net/http"
	"text/template"
	"databases/util"
	"database/sql"
     _ "github.com/go-sql-driver/mysql"
)

var template_html = template.Must(template.ParseGlob("templates/*"))


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


func Home(writer http.ResponseWriter, request *http.Request) {
	var customers []Customer
	customers = GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func Create(writer http.ResponseWriter, request *http.Request) {
	//  var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	template_html.ExecuteTemplate(writer, "Create", nil)

}

func Insert(writer http.ResponseWriter, request *http.Request) {
	//  var customers []Customer
	//customers = GetCustomers()
	//log.Println(customers)
	var customer Customer
	customer.CustomerName = request.FormValue("customername")
	customer.SSN = request.FormValue("ssn")
	InsertCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = request.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	deleteCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)

}

func GetCustomerById(customerId int) Customer {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer WHERE CustomerId=?", customerId)
	if error != nil {
		panic(error.Error())
	}
	//fmt.Println(rows)
	var customer Customer
	customer = Customer{}

	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		error = rows.Scan(&customerId, &customerName, &SSN)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
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
	log.Println("Server started on: http://localhost:8000")
    http.HandleFunc("/", Home)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.ListenAndServe(":8000", nil)

	// var customers []Customer
	// customers = GetCustomers()
	// fmt.Println("Before Update Customers", customers)
	// var customer Customer
	//     customer.CustomerName = "saikiran"
	// customer.SSN = "272432387"
	// customer.CustomerId = 6


	// InsertCustomer(customer)
	// //UpdateCustomer(customer)
	// //deleteCustomer(customer)
	// customers = GetCustomers()
	// fmt.Println("After Update", customers)
}