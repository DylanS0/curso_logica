//Fundamentos de Dart:

//comentario de una linea

/* comentario de varias lineas
primer programa 
suma de dos numeros enteros 
 */

/* Funciones: Bloques de codigo identificado con un 
nombre que realiza una accion, indica el tipo de dato que retorna.
en caso que no retorne nada: void*/

/* Salida de datos: escribir */

/* Tipos de datos: texto: String
numeros: Enteros (int) y Float(double)
logico: Booleanos(true/falso) */

//propiedades: llave
//metodos: cubos

//input: leer: entrada
//output: escribir: salida

//Funcion calsica, anonima, flecha

//contol+click=busca donde esta la funcion

/*Estructuras de ciclos:
for: n veces; se sabe cuantas veces se va a  repetir
while: mientras se valida una condicion; true or false
*/

import './utilidades.dart';

void main() {
  //Declarar variables
  encabezado('Programa para sumar 2 numeros enteros', 130);
  int numero1 = 0;
  int numero2 = 0;
  int numero3 = 0;
  int numero4 = 0;

  numero1 = leerNumeroEntero('=> Ingresa primer numero: ');
  numero2 = leerNumeroEntero('=> Ingresa segundo numero: ');
  numero3 = leerNumeroEntero('=> Ingresa tercer numero: ');
  numero4 = leerNumeroEntero('=> Ingresa cuarto numero: ');

  int suma = calcularSuma([numero1, numero2, numero3, numero4]);
  linea(130);
  print('El resultado de la suma es: $suma ');
  linea(130);
}
