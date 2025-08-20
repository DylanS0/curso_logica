package main
import (
	"fmt"
	"sort"
	ut "module/utilidades"
	"strconv"
	"strings"
)

func main() {
	const nroCiudades int =6 
	var (
		ciudades [] string 
		ciudad string 
	)

	ut.Encabezado("Ingreso de ciudades", 50)
	for i := 0; i < nroCiudades; i++ {
		ciudad = ut.LeerTexto("Ingrese la ciudad " + strconv.Itoa(i+1) )
		ciudades = append(ciudades, ciudad)
	}

	verCiudades(ciudades, "listado de ciudades", 50)

	sort.Strings(ciudades)
	verCiudades(ciudades, "ciudades ordenadas ascendentemente", 50)
	
	sort.Sort(sort.Reverse(sort.StringSlice(ciudades)))
	verCiudades(ciudades, "ciudades ordenadas descendentemente: ", 50)
	
}

func verCiudades ( ciudades []string, titulo string, largo int)  {
	ut.Encabezado(titulo, largo)
	for _, item := range ciudades {
		fmt.Println("La ciudad es: ", strings.ToUpper(item))
	}
}