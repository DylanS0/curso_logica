/* crear un array de ciudades */
import * as util from "./utilidades.js";

const main = () => {
    //definir variables
    const ciudades = [];
    let ciudad = '';
    const avisos = [
        'primera',
        'segunda',
        'tercera', 
        'cuarta',
        'quinta', 
        'sexta'
    ];
    const nroCiudades = 6;

    util.encabezado("ingreso de ciudades", 50);
    for (let i = 0; i < nroCiudades; i++) {
        ciudad = util.leerTexto(`ingresa la ${avisos[i]} ciudad`, 'ciudad');
        ciudades.push(ciudad);
    }

    verCiudades(ciudades, 'mostrar ciudades', 50);

    ciudades.sort();
    verCiudades(ciudades, 'ciudades ordendas ascendentemente', 50);
    
    ciudades.sort((a,b) => b.localeCompare(a));
    verCiudades(ciudades, 'ciudades ordendas descendentemente', 50);
   

}

const verCiudades = (ciudades, titulo = ' ', largo = 0) => {
    util.encabezado(titulo, largo);
    ciudades.forEach((item) => console.log(`la ciudad es ${item.toUpperCase()}`));

}

main()