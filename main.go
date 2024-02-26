package main

import (
	"fmt"
	"sesi-3/model"
)

// func main() {
// 	name := "Gio"
// 	age := 25

// 	greet(name, age)
// }

// func greet(name string, age int) {
// 	fmt.Printf("Hi, I'm %s and I'm %v years old\n", name, age)
// }

// func main() {
// 	fmt.Println(greet("Group 1 members:", []string{"Yoda", "Marsha", "Rini", "Joko"}))
// }

// func greet(msg string, name []string) string {
// 	joinName := strings.Join(name, ", ")

// 	return msg + " " + joinName
// }

// func main() {
// 	diameter := float64(14)
// 	area, circumference := calculateCircle(diameter)

// 	fmt.Printf("Circle width: %v\n", diameter)
// 	fmt.Printf("Circle area: %v\n", area)
// 	fmt.Printf("Circle circumference: %v\n", circumference)
// }

// // This function will return an area and circumference from a circle diameter
// func calculateCircle(diameter float64) (area float64, circumference float64) {
// 	const pi = math.Pi

// 	area = (pi * math.Pow(diameter, 2)) / 4
// 	circumference = pi * diameter

// 	return
// }

// func main() {
// 	studentList := print("Jordan", "Bimo", "Juan", "Giorno")

// 	fmt.Printf("%+v", studentList)
// }

// func print(names ...string) []map[string]string {
// 	result := []map[string]string{}

// 	for i, name := range names {
// 		key := fmt.Sprintf("student-%v", i+1)
// 		result = append(result, map[string]string{
// 			key: name,
// 		})
// 	}

// 	return result
// }

// func main() {
// 	total := sum(5, 5, 2, 3, 5, 1, 2, 2, 10, 5)

// 	fmt.Println("Sum: ", total)
// }

// func sum(numbers ...int) int {
// 	result := 0

// 	for _, num := range numbers {
// 		result += num
// 	}

// 	return result
// }

// func main() {
// 	profile("Joan", "Apple", "Mango", "Grape", "Melon")
// }

// func profile(name string, favouriteFruits ...string) {
// 	favouriteFruit := strings.Join(favouriteFruits, ", ")

// 	fmt.Printf("Hi, my name is %s\n", name)
// 	fmt.Printf("And my favourite fruits are: %s", favouriteFruit)
// }

// func main() {
// 	evenNumbers := func(numbers ...int) []int {
// 		res := []int{}

// 		for _, num := range numbers {
// 			if num%2 == 0 {
// 				res = append(res, num)
// 			}
// 		}

// 		return res
// 	}

// 	fmt.Println(evenNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
// }

// ada

// func main() {
// 	func(str string) {
// 		temp := ""

// 		for i := len(str) - 1; i >= 0; i-- {
// 			temp += string(str[i])
// 		}

// 		fmt.Println("Is palindrome? ", str == temp)
// 	}("123")
// }

// func main() {
// 	studentLists := []string{"Giorno", "Marsha", "Juliet", "Romeo"}

// 	find := findStudent(studentLists)

// 	fmt.Println(find("MARSHA"))
// }

// func findStudent(students []string) func(string) string {
// 	return func(s string) string {
// 		student := ""
// 		position := -1

// 		for i, stud := range students {
// 			if strings.ToLower(s) == strings.ToLower(stud) {
// 				student = stud
// 				position = i
// 			}
// 		}

// 		if position < 0 {
// 			return fmt.Sprintln("Student not found!")
// 		}

// 		return fmt.Sprintf("Found student with name %s in position %v", student, position)
// 	}
// }

// type isOddNumber func(n int) bool

// func main() {
// 	numbers := []int{1, 4, 5, 5, 2}

// 	find := findOddNumbers(numbers, func(numb int) bool {
// 		return numb%2 != 0
// 	})

// 	fmt.Println("Total odd number: ", find)
// }

// func findOddNumbers(numbers []int, callback isOddNumber) int {
// 	totalOddNumber := 0

// 	for _, num := range numbers {
// 		if callback(num) {
// 			totalOddNumber++
// 		}
// 	}

// 	return totalOddNumber
// }

// func main() {
// 	var firstNumber int = 4
// 	var secondNumber *int = &firstNumber

// 	fmt.Println("firstnumber: ", firstNumber)
// 	fmt.Println("firstnumber memory address: ", &firstNumber)
// 	fmt.Println("secondNumber: ", *secondNumber)
// 	// firstNumber = 1
// 	*secondNumber = 1

// 	fmt.Println("firstnumber after update: ", firstNumber)
// 	fmt.Println("secondNumber memory address: ", secondNumber)
// }

// func main() {
// 	number := 10
// 	fmt.Println("Number init: ", number)

// 	change(&number)

// 	fmt.Println("Number after change: ", number)
// }

// func change(num *int) {
// 	*num = 1
// }

// type Person struct {
// 	name string
// 	age  int
// }

// type employeeStatus string

// const (
// 	FulltimeEmployee employeeStatus = "fulltime"
// 	ContractEmployee employeeStatus = "contract"
// )

// func (e *Employee) GetPerson() Person {
// 	return e.person
// }

// func (e *Employee) UpdateToFulltime() {
// 	e.status = FulltimeEmployee
// }

// type Employee struct {
// 	person   Person
// 	division string
// 	status   employeeStatus
// }

// func main() {
// 	employee := Employee{
// 		name:     "John",
// 		age:      35,
// 		division: "Technology",
// 	}
// 	employee.age = 30
// 	employee.division = "Finance"

// 	// fmt.Printf("Employee: %+v", employee)
// 	fmt.Println("Employee name: ", employee.name)
// 	fmt.Println("Employee age: ", employee.age)
// 	fmt.Println("Employee division: ", employee.division)
// }

// func main() {
// 	firstEmployee := Employee{name: "Giorno", age: 25, division: "Marketing"}
// 	secondEmployee := &firstEmployee

// 	fmt.Println("firstEmployee name: ", firstEmployee.name)
// 	fmt.Println("secondEmployee name: ", secondEmployee.name)

// 	secondEmployee.name = "Ruby"

// 	fmt.Println("firstEmployee name: ", firstEmployee.name)
// 	fmt.Println("secondEmployee name: ", secondEmployee.name)
// }

// func main() {
// 	firstEmployee := Employee{
// 		person: Person{
// 			name: "Firza",
// 			age:  40,
// 		},
// 		division: "Technology",
// 	}

// 	// firstEmployee.person.name = "Herald"
// 	// firstEmployee.person.age = 22
// 	// firstEmployee.division = "Finance"

// 	fmt.Printf("%+v\n", firstEmployee)
// }

// func main() {
// 	employee := struct {
// 		person   Person
// 		division string
// 	}{}

// 	employee.person.name = "Loli"
// 	employee.person.age = 21
// 	employee.division = "Business"

// 	fmt.Printf("%+v\n", employee)
// }

// func main() {
// 	employees := []Employee{
// 		{
// 			person: Person{
// 				name: "Marsha",
// 				age:  20,
// 			},
// 			division: "Technology",
// 		},
// 		{
// 			person: Person{
// 				name: "Toni",
// 				age:  25,
// 			},
// 			division: "Finance",
// 		},
// 	}

// 	employees = append(employees, Employee{})

// 	employees[2].person.name = "Bima"
// 	employees[2].person.age = 26
// 	employees[2].division = "Marketing"

// 	for _, employee := range employees {
// 		fmt.Printf("Employee data: %+v\n", employee)
// 	}
// }

// func main() {
// 	employees := []struct {
// 		person   Person
// 		division string
// 	}{
// 		{
// 			person: Person{
// 				name: "Marsha",
// 				age:  20,
// 			},
// 			division: "Technology",
// 		},
// 		{
// 			person: Person{
// 				name: "Toni",
// 				age:  25,
// 			},
// 			division: "Finance",
// 		},
// 	}

// 	employees = append(employees, Employee{})

// 	employees[2].person.name = "Bima"
// 	employees[2].person.age = 26
// 	employees[2].division = "Marketing"

// 	for _, employee := range employees {
// 		fmt.Printf("Employee data: %+v\n", employee)
// 	}
// }

func init() {
	fmt.Println("Hi, from init func!")
}

func main() {
	employee := model.Employee{
		Person: model.Person{
			Name: "Ghian",
			Age:  31,
		},
		Division: "Marketing",
	}

	employee.Greet()
}
