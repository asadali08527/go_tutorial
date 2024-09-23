
# Go Types and Structs Tutorial

This tutorial explains the concepts of types and structs in Go. Structs are used to group related data together, while types allow us to define our own custom types. This tutorial also covers methods for structs, field visibility, and embedded structs.

## Table of Contents

- [What are Types?](#what-are-types)
- [What are Structs?](#what-are-structs)
- [Accessing Struct Fields](#accessing-struct-fields)
- [Field Visibility in Structs](#field-visibility-in-structs)
- [Struct Methods](#struct-methods)
  - [Value Receivers](#value-receivers)
  - [Pointer Receivers](#pointer-receivers)
- [Embedded Structs](#embedded-structs)
- [Example Code](#example-code)

---

## What are Types?

In Go, we can define custom types using the `type` keyword. These types improve the readability and maintainability of our code.

Example:
```go
type Age int
type Name string
```

Custom types allow us to represent real-world entities and abstract details into meaningful names.

---

## What are Structs?

A **struct** in Go is a composite data type that groups together variables under a single unit. Each variable is called a field, and the fields can be of different types.

Example:
```go
type User struct {
    FirstName   string
    LastName    string
    PhoneNumber string
    Age         int
    BirthDate   time.Time
}
```

Structs allow you to represent complex entities like a `User` by organizing related data into a single entity.

### Creating and Initializing Structs

Structs can be initialized in multiple ways:

1. **Zero Value Initialization**:
   ```go
   var user User
   ```

2. **Named Fields Initialization**:
   ```go
   user := User{
       FirstName: "Trevor",
       LastName: "Sawler",
       PhoneNumber: "1 555 555 1212",
   }
   ```

3. **Positional Initialization**:
   ```go
   user := User{"Trevor", "Sawler", "1 555 555 1212", 0}
   ```

---

## Accessing Struct Fields

You can access the fields of a struct using the dot `.` operator.

Example:
```go
user.FirstName = "John"
log.Println(user.FirstName)
```

---

## Field Visibility in Structs

In Go, struct field visibility is determined by the first letter of the field name:

- **Public Fields**: Fields that start with a capital letter (e.g., `FirstName`) are public and can be accessed from other packages.
- **Private Fields**: Fields that start with a lowercase letter (e.g., `age`) are private and cannot be accessed from outside their package.

Example:
```go
type Person struct {
    FirstName string  // Public field
    lastName  string  // Private field
}
```

---

## Struct Methods

Structs in Go can have methods attached to them. These methods add behavior to the data in the struct.

### Value Receivers

A method can have a **value receiver**, which works on a copy of the struct. Changes made in the method do not affect the original struct.

Example:
```go
func (u User) GetFullName() string {
    return u.FirstName + " " + u.LastName
}
```

### Pointer Receivers

A method can have a **pointer receiver**, which allows the method to modify the original struct by dereferencing the pointer.

Example:
```go
func (u *User) SetPhoneNumber(newNumber string) {
    u.PhoneNumber = newNumber
}
```

---

## Embedded Structs

Go does not support inheritance, but you can embed one struct inside another to simulate inheritance-like behavior.

Example:
```go
type Address struct {
    City    string
    State   string
    ZipCode string
}

type Employee struct {
    Name    string
    Address // Embedded struct
}
```

You can access the fields of the embedded struct directly from the parent struct:
```go
emp := Employee{
    Name: "Asad",
    Address: Address{
        City:    "Vienna",
        State:   "Austria",
        ZipCode: "1010",
    },
}
log.Println(emp.City)  // Accessing embedded struct field
```

---

## Example Code

Here is the complete code demonstrating all the concepts of types, structs, and methods in Go:

```go
package main

import (
    "log"
    "time"
)

type User struct {
    FirstName   string
    LastName    string
    PhoneNumber string
    Age         int
    BirthDate   time.Time
}

type Address struct {
    City    string
    State   string
    ZipCode string
}

type Employee struct {
    Name    string
    Address // Embedded struct
}

func (u User) GetFullName() string {
    return u.FirstName + " " + u.LastName
}

func (u *User) SetPhoneNumber(newNumber string) {
    u.PhoneNumber = newNumber
}

func (e Employee) GetAddress() string {
    return e.City + ", " + e.State + " - " + e.ZipCode
}

func main() {
    user := User{
        FirstName:   "Trevor",
        LastName:    "Sawler",
        PhoneNumber: "1 555 555 1212",
        BirthDate:   time.Now(),
    }

    log.Println("User's full name:", user.GetFullName())
    log.Println("User's phone number before change:", user.PhoneNumber)
    user.SetPhoneNumber("1 555 555 1234")
    log.Println("User's phone number after change:", user.PhoneNumber)

    emp := Employee{
        Name: "Asad",
        Address: Address{
            City:    "Vienna",
            State:   "Austria",
            ZipCode: "1010",
        },
    }

    log.Println("Employee's name:", emp.Name)
    log.Println("Employee's address:", emp.GetAddress())
}
```

---

## Conclusion

Structs in Go are powerful tools for organizing and working with complex data. By defining methods on structs, you can add behavior to your types. Embedded structs allow you to share fields between types, creating flexible and reusable code.

