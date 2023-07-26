package main

import "fmt"

//Una profesora de la universidad quiere tener un listado con todos sus estudiantes. Es necesario crear una aplicación que contenga dicha lista.
//Sus estudiantes son: Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
func main() {
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}
	students = append(students, "Juan")
	fmt.Println("Nuevo listado")
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

}
