/* 
calculadora de indice 
*/
import * as util from  "./utilidades.js"

const main = () =>{
    //definir variables
    let peso, altura, imc = 0;

    util.encabezado("calculadora de indice de masa corporal", 50)
    peso = util.leerNumero("ingresa tu peso en kg")
    altura = util.leerNumero("ingresa tu altura en mts")
    imc = calculoImc(peso, altura)
    util.linea(50)
    console.log(`el indice de masa corporal es de ${imc.toFixed(2)} (${categoriasImc(imc).toUpperCase()})`);
    util.linea(50)


}
const calculoImc = (peso = 0, altura = 0) => (peso / Math.pow(altura, 2)) 

const categoriasImc = function(imc = 0){
    if (imc < 18.5) {
        return 'bajo peso'
    } else if (imc >= 18.5 && imc < 24.9) {
        return 'peso normal'
    } else if (imc >= 24.9 && imc < 29.9){
        return 'sobrepeso'
    }else{
        return 'obesidad'
    }

}

main()
