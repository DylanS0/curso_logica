package main

import (
	"fmt" 
	util "module/utilidades"
	"strings"
	"math/rand"
)

func main() {
	//declarar variables
	 var (
		secreto int = rand.Intn(20) +1
		numero  int    = 0
		alias   string = ""
	) 

	util.Encabezado("Juego de adivina el numero", 50)
	alias = util.LeerTexto("Ingrese su alias")
	util.Linea(50)
	jugar(secreto, numero, alias)
}

func jugar(secreto int, numero int, alias string) {
	var intentos int = 1

	for numero != secreto {
		if intentos > 3{
			util.Linea(50)
			fmt.Printf("Perdiste, %v. El numero secreto era %v\n", strings.ToUpper(alias), secreto)
			util.Linea(50)
			break
		}

		numero = util.LeerNumeroEntero("Ingrese su numero")
		intentos++

		if numero > secreto {
			fmt.Printf("%v, El numero es mayor que el secreto...\n", strings.ToUpper(alias))
		} else if numero < secreto{
			fmt.Printf("%v, El numero es menor que el secreto...\n", strings.ToUpper(alias))
		} else{
			fmt.Printf("%v, Has ganado!!!", strings.ToUpper(alias))
		}
	}
} 
