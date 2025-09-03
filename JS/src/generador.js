/* ganerador de tablas de multi */
import { encabezado, leerNumero, linea} from "./utilidades";

const main = () => {
    //definir variables
    let numero = 0
    let inicio = 0
    let final  = 0
    
    encabezado("generador de tablas de multiplicar", 50)
    numero = leerNumero("ingresar tu numero ")
    encabezado(`tabla del ${numero}`, 50)
    generarTablas(numero)
    encabezado("generador de rango de tablas", 50)
    inicio = leerNumero("Ingrese numero de inicio")
    final = leerNumero("ingrese numero final")
    encabezado(`tablas de multiplicar del ${inicio} al ${final}`, 50)
    generarRangoTablas(inicio, final)
}

const generarRangoTablas = function(inicio = 0, final = 0) {
    let nroTablas = Math.abs(inicio - final) +1
    for (let i = 0; i < nroTablas; i++) {
       for (let j = 0; j < 11; j++) {
      console.log(`${inicio} X ${j} = ${producto(inicio, j)}`)
       }
       linea(50)
       inicio++
    }
}


const generarTablas = function (numero = 0) {
    for (let i = 1; i < 11; i++){
        console.log(`${numero} X ${i} = ${producto(numero, i)}`)
    }
}

const producto = (a=0, b=0) => a * b;


main()