package main

import (
	"fmt"
	ut "module/utilidades"
	_ "strings"
)

func main() {
	paisesCapitales := []map[string]string{
        {"pais": "Argentina", "capital": "Buenos Aires"},
        {"pais": "Bolivia", "capital": "Sucre"},
        {"pais": "Brasil", "capital": "Brasilia"},
        {"pais": "Chile", "capital": "Santiago de Chile"},
        {"pais": "Colombia", "capital": "Bogotá"},
        {"pais": "Ecuador", "capital": "Quito"},
        {"pais": "Guyana", "capital": "Georgetown"},
        {"pais": "Paraguay", "capital": "Asunción"},
        {"pais": "Perú", "capital": "Lima"},
        {"pais": "Surinam", "capital": "Paramaribo"},
        {"pais": "Uruguay", "capital": "Montevideo"},
        {"pais": "Venezuela", "capital": "Caracas"},
    }

	/* for _, pais := range paisesCapitales {
		slicePaises = append(slicePaises, pais)
	} */

	/* ut.Encabezado("datos del empleado", 50)
	for clave , valor  := range paisesCapitales {
		fmt.Printf("--> %v: %v\n", strings.ToUpper(clave), strings.ToUpper(valor))
	} */

	ut.Encabezado("Listado de paises de sudamerica", 50)
	for _, pais := range paisesCapitales {
		fmt.Printf("%v: %v\n", pais["pais"], pais["capital"])

	}

}
