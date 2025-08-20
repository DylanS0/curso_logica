/*
funciones de utilidad
sanetizacion

*/
import 'dart:io'; //paquete para input y output

void linea(int largo) {
  print('═' * largo);
}

String centrar(String titulo, int largo) {
  int relleno = ((largo - titulo.length) ~/ 2);
  String textoCentrado = ' ' * relleno + titulo;
  return textoCentrado;
}

void encabezado(String titulo, int largo) {
  linea(largo);
  print(centrar(titulo.toUpperCase(), largo));
  linea(largo);
}

int leerNumeroEntero(String msj) {
  int numero = 0;
  while (true) {
    //inicia ciclo
    try {
      //instrucciones que pueden generar un error
      stdout.write(msj + ': ');
      numero = int.parse(stdin.readLineSync().toString());
      return numero; //rompe el ciclo por el return
    } catch (e) {
      //mensaje de error que maneja el error
      print("Debes escribir un numero");
    }
  }
}

double leerNumeroDecimal(String msj) {
  double numero = 0;

  while (true) {
    try {
      stdout.write(msj + ': ');
      numero = double.parse(stdin.readLineSync().toString());
      return numero;
    } catch (e) {
      print("Debes escribir un numero decimal");
    }
  }
}

int calcularSuma(List<int> numeros) {
  return numeros.reduce(
    (total, numero) => total + numero,
  ); //funcion recursiva, evita ciclos
}

String leerTexto(String msj, String aviso) {
  String texto = '';

  while (true) {
    stdout.write(msj + ': ');
    texto = stdin.readLineSync().toString();

    if (texto.trim().isNotEmpty) { //elimina los espacios en blanco y pregunta si esta vacio. isnotempty = es una propieedad por eso no tiene () en cambio trim es una accion sobre el objeto por eso tiene ()
      return texto.trim();
    } else {
      print('Debes escribir $aviso ...');
    }
  }
}
