package main

import (
	"fmt"
	util "module/utilidades"
	"strconv"
)

func main () {

	util.Encabezado("Generador de tablas de multiplicar", 100)
	numero := util.LeerNumeroEntero("Ingrese su numero ")
	util.Encabezado("Tabla de multiplicar del numero: " + strconv.Itoa(numero)  , 100)
	generarTabla(numero)
	inicio:=util.LeerNumeroEntero("Ingresa el numero de inicio")
	final:=util.LeerNumeroEntero("Ingresa el numero final")
	util.Encabezado("Genera rango de las tablas de multiplicar " + strconv.Itoa(inicio)+ " al " +strconv.Itoa(final), 100)
	rangoTablas(inicio, final)
}

func generarTabla(numero int)  {
	for i := 1; i < 11; i++ {
		fmt.Printf("%v x %v = %v\n", numero, i, producto(numero, i))	
	}
	util.Linea(100)
}

func rangoTablas (inicio, final int)  {
	nroTablas := final - inicio + 1
	for i := 0; i < nroTablas; i++ {
		for j := 1; j < 11; j++ {
			fmt.Printf("%v x %v = %v\n", inicio, j, producto(inicio, j))
		}
		util.Linea(100)
		inicio++
	}
}

func producto (a, b int) int {
	return a * b
}