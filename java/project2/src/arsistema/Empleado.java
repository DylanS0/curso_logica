package arsistema;

import java.util.HashMap;

public class Empleado {
public static void main(String[] args) {
    //declarar un mapa de empleados 
    HashMap<String, String> empleado = new HashMap<>();

    empleado.put("cedula", "12.345.678");
    empleado.put("nombre", "manuel");
    empleado.put("apellido", "lopez");
    empleado.put("edad", "25");
    empleado.put("sueldo", "1234.56$");
    
    Util.Encabezado("datos del empleado", 50);
    for (String clave : empleado.keySet()) {//keySet() muestra clave y values() muestra valores 
        System.out.printf("%s: %s %n", clave.toUpperCase(), empleado.get(clave).toUpperCase());
    }
}
}
