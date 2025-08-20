package arsistema;

public class Indice {
    public static void main(String[] args) {
        //declarar variables
        double peso = 0;
        double altura = 0;
        double imc = 0;
        Util.Encabezado("Indice de masa corporal", 50);
        peso = Util.leerNumeroDecimal("Ingrese su peso");
        altura = Util.leerNumeroDecimal("Ingrese su altura");
        imc = calcularImc(peso, altura);
        Util.linea(50);
        System.out.printf("El indice de masa corporal es de %.2f (%s)  %n", imc, categoriaImc(imc).toUpperCase());
        Util.linea(50);
    }
    public static double calcularImc(double peso, double altura){
        return (peso/ Math.pow(altura, 2));
    }
    public static String categoriaImc(double imc){
        if (imc<18.5) {
            return "Bajo peso";
        } else if (imc >= 18.5 && imc < 24.9 ) {
            return "Peso normal";
        } else if(imc >= 24.9 && imc <= 29.9){
            return "Sobrepeso";
        }else{
            return "Obesidad";
        }
    }
}
