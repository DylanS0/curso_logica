package utilidades
import (
	"fmt"
	"strings"
)

func Linea(largo int) {
	fmt.Println(strings.Repeat("═", largo))
}

func Encabezado (titulo string, largo int){
	Linea(largo)
	fmt.Println(strings.ToUpper((Centrar(titulo, largo))))
	Linea(largo)
}

func Centrar (titulo string, largo int) string {
	centro := ((largo - len(titulo)) / 2)
	tituloCentrado := strings.Repeat(" ", centro) + titulo + strings.Repeat(" ", centro)
	return tituloCentrado
}

func LeerNumeroEntero (msj string) int {
	var numero int
	fmt.Print(msj + ": ")
	fmt.Scanln(&numero)
	return numero
}

func LeerNumeroDecimal (msj string) float64 {
	var numero float64
	fmt.Print(msj + ": ")
	fmt.Scanln(&numero)
	return numero
}

func LeerTexto ( msj string) string {
	var texto string
	fmt.Print(msj + ": ")
	fmt.Scanln(&texto)
	return texto

}