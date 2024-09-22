# Go Variables and Functions Tutorial

This tutorial demonstrates how to declare, initialize, and use variables and functions in Go. It also covers various advanced concepts such as anonymous functions, variadic functions, and multiple return values.

## Table of Contents

- [Variables in Go](#variables-in-go)
  - [Declaring Variables with `var`](#declaring-variables-with-var)
  - [Shorthand Declaration with `:=`](#shorthand-declaration-with-)
  - [Multiple Variable Declaration](#multiple-variable-declaration)
- [Functions in Go](#functions-in-go)
  - [Basic Function Declaration](#basic-function-declaration)
  - [Multiple Return Values](#multiple-return-values)
  - [Variadic Functions](#variadic-functions)
  - [Anonymous Functions](#anonymous-functions)

---

## Variables in Go

### Declaring Variables with `var`

In Go, variables are declared using the `var` keyword, followed by the variable name and its type. Variables declared this way are initialized with their zero values:

- `int`: `0`
- `float64`: `0.0`
- `bool`: `false`
- `string`: `""` (empty string)

Example:

```go
var greeting string  // Declaring a string variable
var number int       // Declaring an integer variable
var decimalNumber float64  // Declaring a floating-point number
var isActive bool    // Declaring a boolean variable
```

### Shorthand Declaration with `:=`

Go also provides a shorthand for declaring and initializing variables in one step. The `:=` operator is used for this purpose and can only be used inside functions. It infers the type of the variable based on the value being assigned.

Example:

```go
city := "Vienna"   // Declares and initializes a string
temperature := 18.5  // Declares and initializes a float64
isOpen := false     // Declares and initializes a boolean
```

### Multiple Variable Declaration

Multiple variables of the same or different types can be declared in one line using commas.

Example:

```go
var country, capital, currency string = "Austria", "Vienna", "Euro"
```

---

## Functions in Go

### Basic Function Declaration

Functions in Go are declared using the `func` keyword. The function takes zero or more parameters and may return zero or more values. Return types must be explicitly declared.

Example:

```go
func calculateArea(length int, width int) int {
    return length * width
}
```

### Multiple Return Values

Go functions can return multiple values, which is a powerful feature. For example, a function can return both the sum and product of two integers:

Example:

```go
func arithmeticOperations(a int, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}
```

### Variadic Functions

A variadic function in Go can accept a variable number of arguments. The last parameter in a variadic function is preceded by an ellipsis (`...`), and it allows you to pass zero or more arguments.

Example:

```go
func sumAll(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
```

### Anonymous Functions

Anonymous functions are functions without names and are typically used for short-lived tasks. They can be assigned to variables or executed immediately.

Example:

```go
twice:= func(x int) int {
    return x * 2
}

fmt.Println(twice(7))  // Outputs: 14
```

Anonymous functions can also be used inline, for example in a print statement:

```go
fmt.Println(func(name string) string {
    return "Hello, " + name
}("Asad"))
```

---

## Conclusion

This tutorial has covered the basics of variables and functions in Go, along with some advanced topics like anonymous functions, variadic functions, and multiple return values. Understanding these concepts is essential for writing efficient Go programs.
