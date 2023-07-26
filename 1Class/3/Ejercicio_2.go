package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Listado normal")
	for _, employee := range employees {
		if employee > 21 {
			fmt.Println(employee)
		}
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Listado act")
	for _, employee := range employees {
		fmt.Println(employee)
	}
}
