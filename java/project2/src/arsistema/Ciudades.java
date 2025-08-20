package arsistema;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

public class Ciudades {
    public static void main(String[] args) {
        // Declarar variables
        List<String> ciudades = new ArrayList<>();
        // String ciudad = new String(); forma vieja
        String ciudad = "";
        final int nroCiudades = 5;

        Util.Encabezado("ingreso de ciudades:", 50);
        for (int i = 0; i < nroCiudades; i++) {
            ciudad = Util.leerTexto("Ingresar la ciudad " + (i + 1), "ciudad");
            ciudades.add(i, ciudad);
        }

        
        verCiudades(ciudades, "ver ciudades", 50);
        ciudades.sort(Comparator.naturalOrder());
        verCiudades(ciudades,"Ver ciudades ordendas ascendentemente:", 50);
        ciudades.sort(Comparator.reverseOrder());
        verCiudades(ciudades, "Ver ciudades ordendas ascendentemente:", 50);
       
    }
    public static void verCiudades(List<String> ciudades, String titulo, int largo){
        Util.Encabezado(titulo, largo);
        ciudades
        .stream()
        .forEach((item) -> System.out.println("La ciudad es " + item.toUpperCase()));
    }
}
