/* estructuras de datos en dart: listas */

import './utilidades.dart';

void main() {
  //declarar una lista de edades
  List<int> edades = [21, 19, 25, 13, 20, 50, 20, 19, 15, 10, 50, 16, 25];
  linea(50);
  print(
    'Cantidad de elementos en la lista: ${edades.length}',
  ); //length longitud de la lista
  linea(50);
  print('La tercera edad es: ${edades[1]}'); //[] indice de la lista
  linea(50);
  print('La ultima edad es ${edades.last}'); // last ultimo elemento de la lista
  encabezado('leer lista de edades con un for clasico', 50);
  for (var i = 0; i < edades.length; i++) {
    //cuando necesito el indice
    print('La edad ${i + 1} es ${edades[i]}');
  }
  encabezado('leer lista de edades con un for in', 50);
  for (var edad in edades) {
    //no me intersa el indice
    print('la edad es $edad');
  }
  encabezado('leer lista de edades con un forEach', 50); //forEach
  edades.forEach((edad) => print('La edad es $edad')); //funcion flecha

  encabezado('mayores de edad', 50);
  edades //encadenar metodos en distinas lineas
      .where((edad) => edad >= 18)
      .forEach((edad) => print('Los mayores de edad $edad')); //ctl+barra

  encabezado('menores de edad', 50);
  edades //encadenar metodos en distinas lineas
      .where((edad) => edad <= 18)
      .forEach((edad) => print('Los menores de edad $edad')); //ctl+barra

  encabezado('ordenar edades ascendentes', 50);
  edades
      .sort(); //sort ordena la lista pero retorna void, asi que no la puedo iterar
  edades.forEach((edad) => print('La edad es ${edad}'));

  encabezado('ordenar edades descendentes', 50);
  edades.sort((a, b) => b - a);
  edades.forEach((edad) {
    print('La edad es $edad');
  });

  encabezado('Edades unicas', 50);
  Set<int> edadesUnicas = edades.toSet();
  //var edadesUnicas = edades.toSet();

  edadesUnicas.forEach((edad) => print('La edad es ${edad}'));
}
