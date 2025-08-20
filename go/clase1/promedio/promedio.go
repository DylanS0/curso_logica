package main

import (
	"fmt"
	ut "module/utilidades"
	"strconv"
)

func main() {
	const nroPersonas int = 6
	var (
		estatura  float64
		sumatoria float64
		promedio  float64
	)
	
	ut.Encabezado("programa para calcular el promedio de estaturas",50)
	
	for i := 0; i < nroPersonas; i++ {
		estatura = ut.LeerNumeroDecimal("Ingrese la estatura #" + strconv.Itoa(i+1) + " en metros ")
		sumatoria += estatura
	}

	promedio = sumatoria / float64(nroPersonas)
	ut.Linea(50)
	fmt.Printf("El promedio de las estaturas es %.2f\n", promedio)
	ut.Linea(50)

}
