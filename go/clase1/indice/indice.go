package main

import (
	"fmt"
	"math"
	ut "module/utilidades"
	"strings"
)
func main() {
	var (
		peso float64
		altura float64
		imc float64
	)
	ut.Encabezado("Calculadora de IMC",100)
	peso = ut.LeerNumeroDecimal("Ingrese peso en kg: ")
	altura = ut.LeerNumeroDecimal("Ingrese altura en mts: ")
	imc = CalcularIMC(peso, altura)
	ut.Linea(100)
	fmt.Printf("El indice de masa corporal es %2f- %v\n", imc, strings.ToUpper(categoria(imc)))
	ut.Linea(100)
	
}

func CalcularIMC(peso float64, altura float64 ) float64 {
	return (peso/(math.Pow(altura, 2)))
}
/* y=&&
o=||
no=!
*/
func categoria(imc float64) string {
	if imc < 18.5 {
		return "Bajo peso"
	} else if imc >= 18.5 && imc < 24.9{
		return "Peso normal"
	} else if imc >=24.9 {
		return "Sobrepeso"
	}else {
		return "Obesidad"
	}
}
