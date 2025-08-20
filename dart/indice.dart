/*
Calculadora de indice de masa corporal
if...else if...else
*/

import 'utilidades.dart' as util;
import 'dart:math' as mth;

void main() {
  //declarar variables
  double peso;
  double altura;
  double imc;

  util.encabezado('Calculadora de indice de masa corporal', 150);

  peso = util.leerNumeroDecimal('Ingresa tu peso en kg');
  altura = util.leerNumeroDecimal('Ingresa tu estatura en mts');
  imc = calculoIMC(peso, altura);

  util.linea(50);
  print('El indice de masa corporal es: ${imc.toStringAsFixed(2)} - ${categoriaImc(imc).toUpperCase()}');
  util.linea(50);
}

double calculoIMC(double peso, double altura) {
  return (peso / mth.pow(altura, 2));
}

String categoriaImc(double imc) {
  if (imc < 18.5) {
    return 'Estas bajo peso';
  } else if (imc >= 18.5 && imc < 24.9) {
    return 'Estas peso normal';
  } else if (imc >= 24.9 && imc < 29.9) {
    return 'Estas sobrepeso';
  } else {
    return 'Estas en obesidad';
  }
}
