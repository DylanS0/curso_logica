package arsistema;

import java.util.Random;

public class Adivina {
    public static Random random = new Random();

    public static void main(String[] args) {
        // declarar variables (atributos de la clase):
        int secreto = random.nextInt(20)+1;
        int numero = 0;
        String alias = " ";

        Util.Encabezado("Bienvenido al juego de Adivina el numero", 50);
        alias = Util.leerTexto("Ingresa tu alias", "alias");
        jugar(secreto, numero, alias);
    }
    public static void jugar (int secreto, int numero, String alias) {
        String mensaje = "";
        int intentos = 1;

        while (numero != secreto) {
            if (intentos == 4) {
                Util.linea(50);
                mensaje = "Perdiste el numero secreto era el: ";
                System.out.printf("%s, %s %d %n", alias.toUpperCase(), mensaje, secreto);
                Util.linea(50);
                break;
            }
            numero = Util.leerNumeroEntero("Ingresa tu numero");
            intentos++;

            if (numero < secreto) {
                mensaje = "tu numero es menor al secreto...";
                System.out.printf("%s, %s %n", alias.toUpperCase(), mensaje);
            } else if(numero > secreto){
                mensaje = "tu numero es mayor al secreto...";
                System.out.printf("%s, %s %n", alias.toUpperCase(), mensaje);
            } else{
                Util.linea(50);
                mensaje = "Felicitaciones ganaste...";
                System.out.printf("%s, %s %n", alias.toUpperCase(), mensaje);
                Util.linea(50);
            }
        }
        
    }
}
