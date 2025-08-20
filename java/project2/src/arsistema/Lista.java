package arsistema;

import java.util.Comparator;
import java.util.HashSet;
import java.util.List;


public class Lista {
    public static void main(String[] args) {
        // Declarar una lista de edades
        // List<Integer> edades = Array.asList(21,19,23,17,21,17,14,15,25,18,15,28);
        // forma antigua
        // int[] edades = new int [12]; forma mas antigua
        // int[] edades = {12,13,13,25,23} forma valida tambien
        List<Integer> edades = List.of(21, 19, 23, 17, 21, 17, 14, 15, 25, 18, 15, 28);
        Util.linea(50);
        System.out.println("Cantidad de edades es: " + edades.size());
        Util.linea(50);
        System.out.println("La tercera edad es: " + edades.get(2));

        Util.Encabezado("Lista de edades con fori:", 50);
        for (int i = 0; i < edades.size(); i++) {
            System.out.printf("-> La edad %d es %d %n", i, edades.get(i));// con el for clasico muestra el indice (muy viejo)
        }
        Util.Encabezado("leer listas de edades con foreach:", 50);

        for (Integer edad : edades) {
            System.out.println("-> La edad es " + edad);// con el for each no muestar el indice (viejo)
        }
        Util.Encabezado("leer lista de datos con streams:", 50);
        edades
        .stream()
        .forEach((edad) -> System.out.println("-> La edad es " + edad));//funcion flecha es lamda

        Util.Encabezado("mayores de edad:", 50);
        edades
        .stream()
        .filter((edad) -> edad >= 18)
        .forEach((edad) -> System.out.println("-> La edad es: " + edad));

        Util.Encabezado("menores de edad:", 50);
        edades
        .stream()
        .filter((edad) -> edad < 18)
        .forEach((edad) -> System.out.println("-> La edad es: " + edad));

        Util.Encabezado("las edades ordenadas ascendente:", 50);
        edades
        .stream()
        .sorted()
        .forEach((edad) -> System.out.println("-> La edad es " + edad));
        
        Util.Encabezado("las edades ordenadas desendentemente:", 50);
        edades
        .stream()
        .sorted(Comparator.reverseOrder())
        .forEach((edad) -> System.out.println("-> La edad es " + edad));

        Util.Encabezado("Listas de edades unicas:", 50);
        HashSet<Integer> edadesUnicas = new HashSet<>(edades);
        edadesUnicas
        .stream()
        .forEach((edad) -> System.out.println("-> La edad es " + edad));
    }
}
