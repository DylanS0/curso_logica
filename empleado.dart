/* 
estructuras de datos: mapas
 */
import './utilidades.dart' as util;
import './paises_suramerica.dart' as paises;

void main() {
  //declarar un mapa de empleado
  Map<String, dynamic> empleado = {
    //dos string pq son clave y error
    'cedula':
        '30.640.501', //determinar si los numeros son string o int, cuando se operan con ellos son int sino string
    'nombre': 'Dylan',
    'apellido': 'Sosa',
    'edad': 20,
    'sueldo': 12000,
  };
  util.encabezado('datos del empleado', 50);
  for (var item in empleado.entries) {
    //entries es para ambos
    //values son los valores
    //keys=llaves del diccionario osea mapa
    print('${item.key.toUpperCase()}: ${item.value}');
  }
  util.encabezado('lista de paises', 50);
  var paisesSur = paises.paisesSurAmerica;
  for (var pais in paisesSur) {
    for (var item in pais.entries) {
      print('${item.key.toUpperCase()}: ${item.value}');
    }
    util.linea(50);
  }
  util.encabezado('Lista de paises ordenados a - z', 50);
  paisesSur.sort(
    (a, b) => a['pais'].toString().compareTo(b['pais'].toString()),
  );

  for (var pais in paisesSur) {
    for (var item in pais.entries) {
      print('${item.key.toUpperCase()}: ${item.value}');
    }
    util.linea(50);
  }
}
