/* greace hopper compilador, go se utiliza para servidores,
package main // Declaración del paquete al que pertenece el código

import "fmt" // Importa el paquete "fmt" para funciones de entrada/salida (como imprimir)
//import "Strings"

func main() { // La función 'main' es donde comienza la ejecución del programa
	suma := sumar(5,3)
	resta := restar(5,3)
	fmt.Println("¡Hola, Mundo!") // Imprime un mensaje en la consola
	fmt.Println("Su suma es: ", suma ) // Imprime un mensaje en la consola
	fmt.Println("Su resta es: ", resta ) // Imprime un mensaje en la consola
}
func sumar (a int,b int) int  {
	return a + b
}
func restar(a float32, b float32) float32 {
	return a - b
}
 func leerSuma() {
	fmt.Println("Ingrese un número a: ")
	fmt.Scanln(&a)
	fmt.Println("Ingrese un número b: ")
	fmt.Scanln(&b)
} ------------------------------------------------------------------------------------------------------------------------------------------------*/

package main

import (
	"fmt"	
	ut "module/utilidades"
	_ "strings"
)

func main() {
	var (
		numero1 int = 0
		numero2 int = 0
		numero3 int = 0
		numero4 int = 0
		suma    int = 0
	)

	ut.Encabezado("sumar numeros", 40)
	//Entrada de datos
	numero1 = ut.LeerNumeroEntero("Ingresa primer numero: ")
	numero2 = ut.LeerNumeroEntero("Ingresa segundo numero: ")
	numero3 = ut.LeerNumeroEntero("Ingresa tercero numero: ")
	numero4 = ut.LeerNumeroEntero("Ingresa cuarto numero: ")

	suma = calcularSuma(numero1, numero2, numero3, numero4)

	ut.Linea(36)
	fmt.Println("La suma de los numeros es : ", suma)
	ut.Linea(36)
}

/*
go mod init module
go mod tidy
para hacer modulos
*/

func calcularSuma(numeros ...int) int { //... convierte al parametro en una lista 
	var sumatoria int = 0
	for _, num := range numeros {
		sumatoria += num
		
	}
	return sumatoria
}
