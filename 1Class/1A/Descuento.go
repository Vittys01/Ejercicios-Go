package main

import "fmt"

// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos.
// Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje.
// La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
func main() {
	fmt.Println(des(100, 10))
}
func des(precio float64, descuento float64) float64 {
	return precio * (1 - descuento/100)
}
