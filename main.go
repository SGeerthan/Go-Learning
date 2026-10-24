package main

import "fmt"

func main() {
    var name string
    var age int
    var height float64
    var isStudent bool

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    fmt.Print("Enter your age: ")
    fmt.Scanln(&age)

    fmt.Print("Enter your height: ")
    fmt.Scanln(&height)

    fmt.Print("Are you a student? (true/false): ")
    fmt.Scanln(&isStudent)

    fmt.Println("\n--- Information ---")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Student:", isStudent)
}