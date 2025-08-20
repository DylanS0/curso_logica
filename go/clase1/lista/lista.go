package main

import (
	"fmt"
	util "module/utilidades"
	"sort"
)


func main() {
	//crear un array de edades
	//edadesArray :=[...] int{15, 20, 10, 25, 18, 22} [...] para no indicar el tamano del arreglo, el compilador lo innfiere
	//crear un slice de edades
	edadesSlice :=[] int{15, 20, 10, 25, 18, 22} 
	
	edadesArray :=[6] int{15, 20, 10, 25, 18, 22} 
	util.Linea(50)
	fmt.Println("--> Edades Array: ", edadesArray)
	fmt.Println("--> Cantidad de edades en array es: ", len(edadesArray)) //len para saber tamano del array 
	//edadesArray = append(edadesArray[:], 18,20,30,45) no se puede agregar un elemento al final del array	pq es array
	util.Linea(50)
	fmt.Println("--> Edades Slice: ", edadesSlice)
	fmt.Println("--> Cantidad de edades en slice es: ", len(edadesSlice))
	edadesSlice = append(edadesSlice, 18,20,30) //agregar un elemento al final del array	
	fmt.Println("--> Lista de edades en slice: ", edadesSlice)
	util.Linea(50)
	fmt.Println("--> Tercera edad en la lista slice de edades: ", edadesSlice[2])
	fmt.Println("--> La ultima edad es: " , edadesSlice[len(edadesSlice)-1])
	fmt.Println("--> La antepenultima edad es: " , edadesSlice[len(edadesSlice)-2])
	util.Linea(50)
	util.Encabezado("--> leer cada edad en la lista slice con un for i: ", 50)
	for i := 0; i < len(edadesSlice); i++ {
		fmt.Printf("La edad %v es: %v\n", i+1, edadesSlice[i]) //%v marcador de posicion 
	}
	util.Encabezado("leer cada edad en la lista slice con un for range: ", 50)
	for _, edad := range edadesSlice {
		//fmt.Printf("La edad %v es: %v\n", i+1, edad)  para mostrar el indice 
		fmt.Println("la edad es : ", edad) 
	}
	util.Encabezado("edades ordenadas ascendentemente", 50)
	sort.Ints(edadesSlice)
	for _, edad := range edadesSlice {
		fmt.Println("Las edades ordenadas son: ", edad)
	}
	util.Encabezado("edades ordenadas descendentemente", 50)
	sort.Sort(sort.Reverse(sort.IntSlice(edadesSlice)))
	for _, edad := range edadesSlice {
		fmt.Println("Las edades ordenadas son: ",edad)
	}

	util.Encabezado("Mayores de edad ", 50)
	for _, edad := range edadesSlice {
		if edad >= 18 {
			fmt.Println("Las edades mayores de edad: ", edad)
		} else {
			fmt.Println("Las edades menores de edad: ", edad)
		}
	}
	/* util.Encabezado("Menores de edad ", 50)
	for _, edad := range edadesSlice {
		if edad < 18 {
			fmt.Println("Las edades menores de edad: ", edad)
		} 
	} */
}