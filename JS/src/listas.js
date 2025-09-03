/* arrays de datos en js  */
import { encabezado, linea } from "./utilidades.js";

const main = () => {
    //definir un array en js
    const edades = [15,49,30,23,23];
    linea(50);
    console.log(`Cantidad de edades es ${edades.length}`);
    linea(50);
    //agregar mas edades 
    edades.push(20);
    edades.push(22);
    edades.push(25);
    console.table(edades);
    edades.push(28);
    edades.push(25);
    linea(50);
    console.log(`la tercera edad es ${edades[2]}`);
    linea(50);
    console.log(`la ultima edad es ${edades.at(-1)}`);
    linea(50);

    encabezado("leer la lista de edades", 50);
    for (const edad of edades) {
        console.log(`la edad es ${edad}`)
    } 

    encabezado("leer la lista de edades (foreach)", 50)    
    edades.forEach((edad) => console.log(`la edad es ${edad}`))

    encabezado("mayores de edad",50);
    edades
    .filter((edad) => edad >= 18)
    .forEach((edad) => console.log(`la edad es ${edad}`));
    
    encabezado("menores de edad",50);
    edades
    .filter((edad) => edad < 18)
    .forEach((edad) => console.log(`la edad es ${edad}`));

    encabezado("edades ordenadas ascendentes",50);
    edades
    .sort()
    .forEach((edad) => console.log(`La edad es ${edad}`));
    
    encabezado("edades ordenadas descendentes",50);
    edades
    .sort((a,b) => b - a )
    .forEach((edad) => console.log(`La edad es ${edad}`));

    encabezado("edades unicas", 50);
    const edadesUnicas = new Set(edades);
    edadesUnicas
    .forEach((edad) => console.log(`la edad es ${edad}`));





}



main()