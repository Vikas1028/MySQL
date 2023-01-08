package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Contacts represents an Contacts in an address book
type Student struct {
	FirstName string
	LastName  string
	Address   string
	City      string
	State     string
	Zip       string
	Phone     string
	Email     string
}

var db *sql.DB

func main() {

	connectToDatabase()
	welcomeMessage()

	var students []Student
	startAddressBook(students)
}

/*
description: to connect with database
*/
func connectToDatabase() {

	var err error
	//"UserName:Password@tcp(portNumber)/databaseName"
	db, err = sql.Open("mysql", "root:Vikas@Sanjay@Bhandekar@28102000@tcp(127.0.0.1:3306)/address_book_service")
	errorHanding(err)

	pingErr := db.Ping()
	errorHanding(pingErr)

}

/*
description: To handle error
param: error
*/
func errorHanding(err error) {
	if err != nil {
		log.Fatalf("Err: ", &err)
	}
}

/*
description: To use address book
param: addressBook []contact
*/
func startAddressBook(students []Student) {

	fmt.Print("\n1)Print Students data \n2)Print Students data by City or State Name \n3)City and State Count \n4)Add Student \n5)Update Student Data  \n6)Delete Student Data \n7)Close Address Book \nChoose: ")
	var task int
	fmt.Scanln(&task)

	switch task {
	case 1:
		//print address Book
		stds, err := listAll(students)
		PrintDatabase(stds, err)
	case 2:
		//Student Data by City Or State
		stds, err := listAllByCitiesOrState(students)
		PrintDatabase(stds, err)
	case 3:
		numberOfContactsByCityOrState()
	case 4:
		//Add Student
		addStudent()
	case 5:
		//Update
		updateStudentData()
	case 6:
		//Delete
		deleteStudentData()
	case 7:
		//close address Book
		fmt.Println("Close Address Book")
		return
	}
	startAddressBook(students)
}

/*
description: welcomeMessage prints a welcome message to the console
*/
func welcomeMessage() {
	fmt.Println("Welcome to Address Book")
}

/*
description: Retrive all data from the database
param: slice of struct(students)
return: slice of struct(students) and error
*/
//UC9
func listAll(students []Student) ([]Student, error) {

	rows, err := db.Query("SELECT * FROM address_book_table;")
	errorHanding(err)
	defer rows.Close()

	// Loop through rows, using Scan to assign record to slice
	for rows.Next() {
		var std Student

		if err := rows.Scan(&std.FirstName, &std.LastName, &std.Address, &std.City, &std.State, &std.Zip, &std.Phone, &std.Email); err != nil {
			return nil, fmt.Errorf("error in query all student: %v", err)
		}
		students = append(students, std)
	}
	return students, nil
}

/*
desription: Print retrived data line by line
param: students data and error
*/
func PrintDatabase(stds []Student, err error) {

	errorHanding(err)
	for _, data := range stds {
		fmt.Println(data)
	}
}

/*
description: Retrive all data from the database By using City Or State Name
param: slice of struct(students)
return: slice of struct(students) and error
*/
//UC10
func listAllByCitiesOrState(students []Student) ([]Student, error) {

	var CityOrState string
	fmt.Print("Enter City Or State name: ")
	fmt.Scanln(&CityOrState)
	rows, err := db.Query("SELECT * FROM address_book_table where City = '" + CityOrState + "' or State = '" + CityOrState + "';")
	errorHanding(err)
	defer rows.Close()

	// Loop through rows, using Scan to assign record to slice
	for rows.Next() {
		var std Student

		if err := rows.Scan(&std.FirstName, &std.LastName, &std.Address, &std.City, &std.State, &std.Zip, &std.Phone, &std.Email); err != nil {
			return nil, fmt.Errorf("error in query all student: %v", err)
		}
		students = append(students, std)
	}
	return students, nil
}

/*
description: Count Number of Contacts in the Database by City or State
*/
// UC11
func numberOfContactsByCityOrState() {

	var count int
	var location string

	fmt.Print("Enter City or State name: ")
	fmt.Scanln(&location)

	err := db.QueryRow("SELECT COUNT(*) FROM address_book_table WHERE City = '" + location + "' OR State = '" + location + "';").Scan(&count)
	errorHanding(err)
	fmt.Println(location, ": ", count)
}

/*
description: Add new student data
*/
// UC12
func addStudent() {

	var FirstName, LastName, Address, City, State, Email string
	var Zip, Phone int
	fmt.Print("Enter first name: ")
	fmt.Scanln(&FirstName)
	fmt.Print("Enter last name: ")
	fmt.Scanln(&LastName)
	fmt.Print("Enter address: ")
	fmt.Scanln(&Address)
	fmt.Print("Enter city: ")
	fmt.Scanln(&City)
	fmt.Print("Enter state: ")
	fmt.Scanln(&State)
	fmt.Print("Enter zip code: ")
	fmt.Scanln(&Zip)
	fmt.Print("Enter phone number: ")
	fmt.Scanln(&Phone)
	fmt.Print("Enter email: ")
	fmt.Scanln(&Email)

	_, err := db.Exec("INSERT INTO address_book_table (First_name, Last_name, Address, City, State, Zip, Phone_number, Email) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", FirstName, LastName, Address, City, State, Zip, Phone, Email)
	errorHanding(err)
	fmt.Println("New Student Data Added: ")
}

/*
description: Update students data
*/
// UC14
func updateStudentData() {

	field := []string{"First_name", "Last_name", "Address", "City", "State", "Zip", "Phone_number", "Email"}
	fmt.Println("Which Field do you want to Update: ")
	fmt.Print("1)FirstName \n2)Last_name \n3)Address \n4)City \n5)State \n6)Zip \n7)Phone_number \n8)Email \nChoose: ")
	var task int
	fmt.Scanln(&task)
	var data string
	switch task {
	case 1:
		fmt.Print("Enter first name: ")
		fmt.Scanln(&data)
	case 2:
		fmt.Print("Enter last name: ")
		fmt.Scanln(&data)
	case 3:
		fmt.Print("Enter address: ")
		fmt.Scanln(&data)
	case 4:
		fmt.Print("Enter city: ")
		fmt.Scanln(&data)
	case 5:
		fmt.Print("Enter state: ")
		fmt.Scanln(&data)
	case 6:
		fmt.Print("Enter zip code: ")
		fmt.Scanln(&data)
	case 7:
		fmt.Print("Enter phone number: ")
		fmt.Scanln(&data)
	case 8:
		fmt.Print("Enter email: ")
		fmt.Scanln(&data)
	}

	fmt.Println("By Using Which Field Condition do you want to Update: ")
	fmt.Print("1)FirstName \n2)Last_name \n3)Address \n4)City \n5)State \n6)Zip \n7)Phone_number \n8)Email \nChoose: ")
	var task2 int
	fmt.Scanln(&task2)
	var data2 string
	switch task2 {
	case 1:
		fmt.Print("Enter first name: ")
		fmt.Scanln(&data2)
	case 2:
		fmt.Print("Enter last name: ")
		fmt.Scanln(&data2)
	case 3:
		fmt.Print("Enter address: ")
		fmt.Scanln(&data2)
	case 4:
		fmt.Print("Enter city: ")
		fmt.Scanln(&data2)
	case 5:
		fmt.Print("Enter state: ")
		fmt.Scanln(&data2)
	case 6:
		fmt.Print("Enter zip code: ")
		fmt.Scanln(&data2)
	case 7:
		fmt.Print("Enter phone number: ")
		fmt.Scanln(&data2)
	case 8:
		fmt.Print("Enter email: ")
		fmt.Scanln(&data2)
	}

	_, err := db.Exec("UPDATE address_book_table SET " + field[task-1] + " = '" + data + "' WHERE " + field[task2-1] + " = '" + data2 + "';")
	errorHanding(err)

	fmt.Println("Student Data Update Successfully")
}

/*
description: Delete students data
*/
// UC14
func deleteStudentData() {

	var err error
	fmt.Println("Delete Student Data By Which Field: ")
	fmt.Print("1)FirstName \n2)Last_name \n3)Address \n4)City \n5)State \n6)Zip \n7)Phone_number \n8)Email \nChoose: ")
	var task int
	fmt.Scanln(&task)
	var data string
	switch task {
	case 1:
		fmt.Print("Enter first name: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where First_name = ?", data)
	case 2:
		fmt.Print("Enter last name: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where Last_name = ?", data)
	case 3:
		fmt.Print("Enter address: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where Address = ?", data)
	case 4:
		fmt.Print("Enter city: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where City = ?", data)
	case 5:
		fmt.Print("Enter state: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where State = ?", data)
	case 6:
		fmt.Print("Enter zip code: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where Zip = ?", data)
	case 7:
		fmt.Print("Enter phone number: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where Phone_number = ?", data)
	case 8:
		fmt.Print("Enter email: ")
		fmt.Scanln(&data)
		_, err = db.Exec("DELETE FROM address_book_table where Email = ?", data)
	}

	errorHanding(err)
	fmt.Println("Student Data deleted Successfully")
}
