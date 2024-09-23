
// Package declaration
package main

import (
    "log"
    "time"
)

// *** Struct Definition ***

// User struct: This struct represents a user entity with fields for storing personal information.
// Struct fields starting with capital letters are exported (public), meaning they can be accessed from outside the package.
// Fields with lowercase letters would be private and accessible only within the same package.
type User struct {
    FirstName   string    // Public field
    LastName    string    // Public field
    PhoneNumber string    // Public field
    Age         int       // Public field
    BirthDate   time.Time // Public field
}

// Address struct: This struct represents a postal address.
type Address struct {
    City    string
    State   string
    ZipCode string
}

// Employee struct: Embedding the Address struct inside the Employee struct.
// This allows Employee to "inherit" fields from Address, which can be accessed directly from an Employee instance.
type Employee struct {
    Name    string
    Address // Embedded struct
}

// *** Struct Methods ***

// Method with a value receiver: This method does not modify the original struct because it works on a copy.
func (u User) GetFullName() string {
    return u.FirstName + " " + u.LastName
}

// Method with a pointer receiver: This method modifies the original struct.
func (u *User) SetPhoneNumber(newNumber string) {
    u.PhoneNumber = newNumber
}

// Method for the Employee struct that returns the full address.
func (e Employee) GetAddress() string {
    return e.City + ", " + e.State + " - " + e.ZipCode
}

// Main function: Entry point of the program
func main() {
    // *** Struct Initialization ***

    // Initializing a User struct with named fields
    user := User{
        FirstName:   "Trevor",
        LastName:    "Sawler",
        PhoneNumber: "1 555 555 1212",
        BirthDate:   time.Now(), // Default value for BirthDate as current time
    }

    // Logging initial details of the user
    log.Println("User's full name:", user.GetFullName())
    log.Println("User's phone number before change:", user.PhoneNumber)

    // Calling method with pointer receiver to change the phone number
    user.SetPhoneNumber("1 555 555 1234")

    // Logging updated phone number
    log.Println("User's phone number after change:", user.PhoneNumber)

    // Initializing an Employee struct with an embedded Address struct
    emp := Employee{
        Name: "Asad",
        Address: Address{
            City:    "Vienna",
            State:   "Austria",
            ZipCode: "1010",
        },
    }

    // Accessing fields from the embedded Address struct
    log.Println("Employee's name:", emp.Name)
    log.Println("Employee's address:", emp.GetAddress())

    // *** Struct Field Visibility ***

    // Demonstrating the visibility of struct fields:
    // Fields starting with uppercase letters (like FirstName) are public and accessible from outside the package.
    // Fields starting with lowercase letters would be private and not accessible.
    myVar := myStruct{
        FirstName: "John", // Public field
    }

    log.Println("myVar's First Name:", myVar.printFirstName())

    myVar2 := myStruct{
        FirstName: "Mary",
    }
    log.Println("myVar2's First Name:", myVar2.printFirstName())
}

// *** Another Struct with Methods ***

// myStruct struct definition with a public field
type myStruct struct {
    FirstName string
}

// Method attached to myStruct to return the first name
func (m myStruct) printFirstName() string {
    return m.FirstName
}
