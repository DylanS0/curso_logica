/*
primer programa en js
que sume 2 numeros
*/

import { encabezado, linea, leerNumero } from "./utilidades.js";

const main = function(){
    let numero1 = 9;
    let numero2 = 8;
    let numero3 = 8;
    let numero4 = 8;
    let suma = 0;

   encabezado('sumar numeros', 50);
   //entradad de datos desde el teclado
   numero1 = leerNumero('ingrese el primer numero');
   numero2 = leerNumero('ingrese el segundo numero');
   numero3 = leerNumero('ingrese el tercer numero');
   numero4 = leerNumero('ingrese el cuarto numero');
   suma = calcularSuma(numero1, numero2);
   console.log("la suma es", suma);
   linea(50)
}
//operador spread
const calcularSuma = function (...numeros) {
    return numeros.reduce((total,numero) => total + numero , 0);
}

main()
//https://github.com/DylanS0/curso_logica.git