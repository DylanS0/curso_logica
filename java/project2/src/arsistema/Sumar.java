package arsistema;//ctl + . para hacer carpeta y meter la clase en la misma
//metodos = acciones
//propiedades = caracteristicas

public class Sumar {
    public static void main(String[] args) throws Exception {
        //declarar variables
        int num1 = 0;
        int num2 = 0;
        int num3 = 0;
        int num4 = 0;
        
        int suma = 0;//buena practica inicializa en 0
        Util.Encabezado("Sumar numeros", 50);

        //entrada de datos
        num1 = Util.leerNumeroEntero("Ingresa numero 1");
        num2 = Util.leerNumeroEntero("Ingresa numero 2");
        num3 = Util.leerNumeroEntero("Ingresa numero 3");
        num4 = Util.leerNumeroEntero("Ingresa numero 4");


        suma = Util.calcularSuma(num1, num2, num3, num4);
        System.out.println("La suma es : " + suma);
        Util.linea(50);
    
    
    }
}
