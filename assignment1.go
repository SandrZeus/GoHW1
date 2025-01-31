package main

import (
	"fmt"
	"errors" //for Task 6 to not get an error
)

func main() {
    //for Task 1
	HelloWorld() 

    //for Task 2
	VariablesAndDataTypes()

    //for Task 3
	ControlStructures()

    //for Task 4
	num := 34
    fact := Functions(num)
    fmt.Printf("Factorial of %d: %d\n", num, fact)

    //for Task 5
	SlicesAndMaps()

    //for Task 6
	a := 12.5
    b := 0.0
    result, err := ErrorHandling(a, b)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Division result: %.2f\n", result)
    }

    //for Task 7
	book := Book{ //assign book details
        Title:           "The Metamorphosis",
        Author:          "Franz Kafka",
        PublicationYear: 1915,
    }
    book.printDetails()
}

//Task 1

func HelloWorld() {
	fmt.Println("Hello, World!")
}

//Task 2

func VariablesAndDataTypes() {
	var (
		anInteger int = 7
		aString string = "Sandro"
		aFloat float64 = 8.6
	)
	fmt.Printf("Integer: %d, Type: %T\n", anInteger, anInteger)
	fmt.Printf("String: %s, Type: %T\n", aString, aString)
    fmt.Printf("Float: %f, Type: %T\n", aFloat, aFloat)
}

//Task 3

func ControlStructures() {
	var input int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&input) //add an input value

	if input%2 == 0{
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	} //declare wheather it's odd or even

	for i :=1; i <= input; i++{
		fmt.Println(i)
	} //count from 1 to the input value
}

//Task 4

func Functions(n int) int {
	if n == 0 {
		return 1
	}
	return n * Functions(n-1)
}

//Task 5

func SlicesAndMaps() {
    
    numbers := []int{1, 2, 3, 4, 5} // Create a slice of integers

    sum := 0
    for _, num := range numbers {
        sum += num
    } // Calculate the sum of elements in the slice

    fmt.Println("Sum of elements in the slice:", sum)

    dictionary := map[string]int{
        "something1":  10,
        "something2": 5,
    } // Create a map

    dictionary["something3"] = 7
	fmt.Println("Value for something1: ", dictionary["something1"])
    fmt.Println("Value for something2: ", dictionary["something2"])
	fmt.Println("Value for something3: ", dictionary["something3"]) // Add and retrieve values from the map
}

//Task 6

func ErrorHandling(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero") // If the division is zero, return an error with a message
    }
    return a / b, nil
}

//Task 7

type Book struct {
    Title string
    Author string
    PublicationYear int
}

func (b Book) printDetails() {
    fmt.Printf("Title: %s\nAuthor: %s\nPublication Year: %d\n", b.Title, b.Author, b.PublicationYear)
}