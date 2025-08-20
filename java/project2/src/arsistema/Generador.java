package arsistema;

public class Generador {
    public static void main(String[] args) {
        // variables
        int numero = 0;
        int inicio = 0;
        int fin = 0;
        Util.Encabezado("generador de tablas de multiplicar", 50);
        numero = Util.leerNumeroEntero("Ingresa tu numero");
        Util.Encabezado("tabla de multiplicar del " + numero, 50);
        tablas(numero);
        Util.Encabezado("generar rango de tablas", 50);
        inicio = Util.leerNumeroEntero("ingresa el inicio del rango");
        fin = Util.leerNumeroEntero("ingresa el fin del rango");
        Util.Encabezado("tablas de multi[plicar del " + inicio + "al " + fin, 50);
        rangoTablas(inicio, fin);
    }
    private static void rangoTablas(int inicio, int fin) {
        int nroTablas = Math.abs(inicio - fin) +1 ;
        for (int i = 0; i < nroTablas; i++) {
            for (int j = 1; j < 11; j++) {
                System.out.printf("%d X %d = %d %n", inicio, j, producto(inicio, j));
            }
            Util.linea(50);
            inicio++;
        }
    }
    public static void tablas (int numero){
        for (int i = 1; i < 11; i++) {
            System.out.printf("%d X %d = %d %n", numero, i, producto(numero, i));
        }
    }

    public static int producto(int a, int b){
        return a * b;
    }

}
