//adivina el numero

import 'dart:math' as math;
import './utilidades.dart';

void main() {
  //declarar variables
  int secreto = math.Random().nextInt(10) + 1;
  int numero = 0;
  String alias = '';

  encabezado('Adivina el numero', 50);
  alias = leerTexto('Ingresa un alias', 'alias');
  linea(50);

  jugar(secreto: secreto, numero: numero, alias: alias);
}

void jugar({required secreto, required numero, required alias}) {
  int intentos = 1;
  while (numero != secreto) {
    if (intentos > 3) {
      linea(50);
      print(
        '${alias.toUpperCase()}, perdiste le numero secreto era el ${secreto}',
      );
      linea(50);
      break;
    }
    intentos++;
    numero = leerNumeroEntero('Ingresa tu numero: ');
    if (numero < secreto) {
      print('tu numero es menor que el secreto...');
    } else if (numero > secreto) {
      print('tu numero es mayor que el secreto...');
    } else {
      linea(50);
      print(' ${alias.toUpperCase()}, Felicidades ganaste...');
      linea(50);
    }
  }
}
