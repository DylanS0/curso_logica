"""diccionario de datos de empleados"""
from utilidades  import *
from typing import Dict, Any

def main():
    #definir diccionario
    empleado: Dict[str, Any] = {
        'ci': '30640501',
        'nombre': 'Juan Carlos',
        'edad': 25,
        'cargo': 'Desarrollador',
        'salario': '5000',
    }
    encabezado('datos del empleado', 50)
    for clave, valor in empleado.items():
        if isinstance(valor, int ):
            print(f"{clave.upper()}: {valor} años")
        else: 
            print(f"{clave.upper()}: {valor}")
    linea(50)
        
main()
