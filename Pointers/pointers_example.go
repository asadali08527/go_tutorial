
// Package declaration
package main

import "fmt"

// Main function: entry point of the program
func main() {
    // Basic Pointer Example

    // Step 1: Declare a string variable and initialize it
    var color string
    color = "Blue"
    
    // Step 2: Print the initial value of the variable
    fmt.Println("Initial value of color:", color)

    // Step 3: Call the changeColor function and pass the memory address of 'color' using the '&' operator
    changeColor(&color)

    // Step 4: Print the value of 'color' after the function call. The value will be changed.
    fmt.Println("Value of color after function call:", color)

    // Pointer with an Integer Example

    // Step 1: Declare an integer variable
    num := 42

    // Step 2: Pass the memory address of 'num' to the increment function
    increment(&num)

    // Step 3: Print the value of 'num' after the function call
    fmt.Println("Value of num after increment:", num)

    // Nil Pointer Example

    // Step 1: Declare a pointer to an int and initialize it with 'nil'
    var ptr *int

    // Step 2: Check if the pointer is nil and print the result
    if ptr != nil {
        fmt.Println("Pointer is not nil, it points to:", *ptr)
    } else {
        fmt.Println("Pointer is nil")
    }
}

// Function: changeColor
// This function takes a pointer to a string as an argument and changes the value of the string.
// Parameters:
// - c: A pointer to a string (*string)
func changeColor(c *string) {
    // Dereference the pointer to modify the value it points to
    *c = "Orange"
}

// Function: increment
// This function takes a pointer to an int and increments its value by 1.
// Parameters:
// - n: A pointer to an int (*int)
func increment(n *int) {
    // Dereference the pointer to increment the value
    *n = *n + 1
}
