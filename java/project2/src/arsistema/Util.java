package arsistema;

import java.util.Arrays;
import java.util.Scanner;

public class Util {
    public static Scanner teclado = new Scanner(System.in);

    public static void linea(int largo) {
        System.out.println("═".repeat(largo));
    }

    public static String centrar(String texto, int largo) {
        int relleno = ((largo - texto.length()) / 2);
        String textoCentrado = " ".repeat(relleno) + texto;
        return textoCentrado;
    }

    public static void Encabezado(String titulo, int largo) {
        linea(largo);
        System.out.println(centrar(titulo.toUpperCase(), largo));
        linea(largo);
    }

    public static int leerNumeroEntero(String msj) {
        int numero = 0;
        while (true) {
            try {

                System.out.print(msj + ": ");
                numero = teclado.nextInt();
                return numero;

            } catch (Exception e) {
                System.out.println("Debe escribir un numero...");
                teclado.next();
            }
        }

    }

    public static int calcularSuma(int... numeros) {
        return Arrays.stream(numeros).sum();
        /*
         * return Arrays.stream(numeros).reduce(0,(total, num) -> total + num);
         * int sumatoria = 0;
         * for (int i = 0; i < numeros.length; i++) {
         * sumatoria += numeros[i];
         * }
         * return sumatoria;
         */
    }

    public static double leerNumeroDecimal(String msj) {
        double numero = 0;

        while (true) {
            try {
                System.out.print(msj + ": ");
                numero = teclado.nextDouble();
                return numero;
            } catch (Exception e) {
                System.out.println("Debe ingresar un numero decimal...");
                teclado.next();
            }
        }
    }

    public static String leerTexto(String msj, String aviso) {
        String texto = "";
        while (true) {
            System.out.print(msj + ": ");
            texto = teclado.nextLine();
            if (!texto.trim().isEmpty()) {
                return texto.trim();
            } else {
                System.out.printf("Debe escribir tu %s...%n", aviso);
                
            }

        }
    }
}
