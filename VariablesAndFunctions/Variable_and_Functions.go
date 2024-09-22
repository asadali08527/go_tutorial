// Package declaration
package main

// Importing necessary packages
import "fmt"

// main function: entry point of the Go program
func main() {
	fmt.Println("Welcome to the Go tutorial on Variables and Functions!")

	// Variables in Go are declared using the 'var' keyword followed by the variable name and its type.
	// Explicit type declaration
	var greeting string       // Declaring a string variable
	var number int            // Declaring an integer variable
	var decimalNumber float64 // Declaring a floating-point number
	var isActive bool         // Declaring a boolean variable

	// Variables can be initialized separately using the '=' operator after declaration.
	greeting = "Hello, Go World!"
	number = 42
	decimalNumber = 3.1415
	isActive = true

	// Printing the values of variables
	fmt.Println("Greeting:", greeting)
	fmt.Println("Number:", number)
	fmt.Println("Decimal Number:", decimalNumber)
	fmt.Println("Is Active:", isActive)

	// Variables can also be declared and initialized at the same time using the shorthand ':=' syntax.
	// This form infers the type based on the value provided and can only be used inside functions.
	city := "Vienna"    // Declares and initializes a string
	temperature := 18.5 // Declares and initializes a float64
	isOpen := false     // Declares and initializes a boolean

	// Printing shorthand initialized variables
	fmt.Println("\nUsing shorthand declaration:")
	fmt.Println("City:", city)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Is Open:", isOpen)

	// Multiple variables can be declared and initialized in one line
	var country, capital, currency string = "Austria", "Vienna", "Euro"
	fmt.Println("\nMultiple variable declaration and initialization:")
	fmt.Println("Country:", country, ", Capital:", capital, ", Currency:", currency)

	// Calling functions and handling return values
	firstName, lastName := getFullName()
	fmt.Println("\nFull Name:", firstName, lastName)

	area := calculateArea(5, 10)
	fmt.Println("Area of the rectangle:", area)

	sum, product := arithmeticOperations(5, 3)
	fmt.Println("Sum and Product of 5 and 3:", sum, product)

	// Variadic function example
	totalSum := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum of all numbers:", totalSum)

	// Anonymous function example
	message := func(name string) string {
		return "Welcome, " + name
	}
	fmt.Println(message("Asad"))

	// **Anonymous Functions**:
	// Functions in Go can also be defined without a name and assigned to a variable.
	// These are known as anonymous functions and are useful for short, one-time tasks.
	twice := func(x int) int {
		return x * 2
	}

	// Using the anonymous function
	result := twice(7)
	fmt.Println("\nThe double of 7 is:", result)

	// Another example of anonymous function used directly in a print statement
	fmt.Println("Anonymous function with string:", func(name string) string {
		return "Hello, " + name
	}("Asad"))

}

// **Functions in Go**:
// Functions are declared using the 'func' keyword followed by the function name.
// The parameters are listed inside parentheses, followed by the return type (if any).

// Function returning multiple values
func getFullName() (string, string) {
	return "Asad", "Ali"
}

// A function can take one or more arguments, and it may or may not return a value.
// The type of each parameter must be explicitly declared, and the return type is specified after the parameters.
func calculateArea(length int, width int) int {
	return length * width
}

// Go allows multiple return values from a function.
// Here, we return both the sum and product of two integers.
func arithmeticOperations(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

// **Variadic Functions**:
// Go supports variadic functions that can accept a variable number of arguments.
// The '...' before the type indicates that the function can take any number of arguments of that type.
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
