package arsistema;

public class Promedio {
    public static void main(String[] args) {
        // Declarar variables
        double altura, promedio, sumatoria = 0;
        final int numeroPersonas = 6;
        double maximo = 0;
        double minimo = 0;
        Util.Encabezado("Promedio de las alturas en cm: ", 50);

        for (int i = 0; i < numeroPersonas; i++) {
            altura = Util.leerNumeroDecimal("Ingresa la altura " + (i+1) );
            sumatoria += altura;
            if (i == 0) { // Primera altura
                maximo = altura;
                minimo = altura;
            } 
            if (altura > maximo) {//ordena el maximo
                maximo = altura;
            }    
            if (altura < minimo) {//ordena el minimo
                minimo = altura;
            }




        }
        promedio = sumatoria / numeroPersonas;
        
        Util.linea(50);
        System.out.printf("El promedio de las alturas es %.2f en cm %n", promedio);
        Util.linea(50);
        System.out.printf("La maxima altura es %.2f en metros %n", maximo);
        Util.linea(50);
        System.out.printf("La minima altura es %.2f en metros %n", minimo);
        Util.linea(50);
    }
}
