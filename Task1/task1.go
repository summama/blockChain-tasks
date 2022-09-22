package main

import "fmt"

type person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(per person) person {
	fmt.Println("Name: ", per.name)
	fmt.Println("Age: ", per.age)
	fmt.Println("Job: ", per.job)
	fmt.Println("Salary: ", per.salary)
	fmt.Println("\n ")
	return per
}

func main() {
	var per1 person
	var per2 person

	// Pers1 specification
	per1.name = "Smmama"
	per1.age = 20
	per1.job = "Student"
	per1.salary = 0
	// Pers2 specification
	per2.name = "Ahmed"
	per2.age = 50
	per2.job = "Teacher"
	per2.salary = 4500
	// print person1 details
	printPerson(per1)
	// print person2 details
	printPerson(per2)

}
