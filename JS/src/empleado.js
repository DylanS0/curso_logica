/* datos del empleado */
import * as util  from "./utilidades.js";

const main = () => {
//declara variables crear un objetivo de js
const empleado = {
    'cedula': '30640501',
    'nombre': 'dylan',
    'apellido': 'Sosa',
    'edad': 20,
    'sueldo': 20000
}

util.encabezado('datos del empleado', 50);
for (const clave in empleado) {
   console.log(`${clave.toUpperCase()}: ${empleado[clave]}`)
        
    }
    //conveertir objeto a un json (texto plano)
    util.linea(50)
    const empleadoJSON = JSON.stringify(empleado);
    console.log(`JSON: empleadoJSON`);

    //convertir un json a un objeto js
    util.linea(50)
    const empleadoObject = JSON.parse(empleadoJSON)
    for (const clave in empleadoObject) {
        console.log(`${clave.toUpperCase()}: ${empleadoObject[clave]}`)
            
        }
}


main()