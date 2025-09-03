/* 
funciones de utilidad
 */
import promptSync from "prompt-sync";
const prompt = promptSync();

const center = function(texto='', largo=0){
    let relleno = (largo - texto.length)/2 ;
    let textoCentrado = ' '.repeat(relleno) + texto;
    return textoCentrado;
} 

const linea = function(largo=0) {
    console.log('═'.repeat(largo));
}

const encabezado = function(titulo='', largo=0){
    linea(largo);
    console.log(center(titulo.toUpperCase(),largo));
    linea(largo);
}

const leerNumero = function (msj = '') {
    let numero = 0;
    while (true) {
        numero = Number(prompt(msj + ': '));
        if (!isNaN(numero) && isFinite(numero)) {
            return numero            
        } else {
            console.log("debe escribir un numero...")
        }
    }
}

const leerTexto = function (msj = '', aviso = '') {
    let texto = ''
    while (true) {
        texto = prompt(msj + ": ")
        if (texto.trim()) {
            return texto.trim()
        }else{
            console.log(`debe escribir ${aviso}...`)
        }
    }
}

export {
    encabezado,
    linea,
    leerNumero, 
    leerTexto

}