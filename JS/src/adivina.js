/* juego de adivina el numero */
import * as util from "./utilidades.js";

const main = () => {
    //definir variables
    let secreto = Math.floor(Math.random()* 20) +1;
    let numero = 0;
    let alias = '';
   util.encabezado("juego de adivina el numero", 50)
    alias = util.leerTexto('ingresa tu alias', 'alias')
    //console.log({alias})
    util.linea(50)
    jugar(secreto, numero, alias) 
   // console.log({secreto})
}

const jugar = function (secreto = 0, numero = 0, alias= '') {
    let msj = ''
    let  intentos = 1;
    while (numero != secreto) {
        if (intentos == 4){
            util.linea(50)
            msj = 'perdiste el numero secretom era '
            console.log(`${alias.toUpperCase()}, ${msj} ${secreto}`)
            util.linea(50)
            break
        }
        numero = util.leerNumero('ingresa tu numero')
        intentos++

        if (numero < secreto){
            msj = 'tu numero es menor que el secreto...'
            console.log(`${alias.toUpperCase()}, ${msj}`)
        } else if(numero> secreto) {
             msj = 'tu numero es mayor que el secreto...'
            console.log(`${alias.toUpperCase()}, ${msj}`)
        }else{
            msj = 'felicidades ganaste'
            util.linea(50)
            console.log(`${alias.toUpperCase()}, ${msj}`)
        }

    }
}

main()