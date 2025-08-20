"""agregar lista de ciudades """
from utilidades import *
from typing import List, Tuple

def main():
    #definir vaeriables
    ciudades: List[str] = []
    ciudad: str 
    msj: Tuple[str, ...] = ("primera","segunda", "tercera", "cuarta", "quinta", "sexta") 
    nro_ciudades: int = 6
    encabezado("ingreso de ciudades", 50)
    for i in range(nro_ciudades):
        ciudad = leerTexto(f"ingrese la {msj[i]} ciudad", "ciudad")
        ciudades.append(ciudad)
        
    mostrar_ciudades(ciudades, 'ver ciudades', 50)
    
    ciudades_asc = list(sorted(ciudades))
    mostrar_ciudades(ciudades_asc, 'ciudades ordenadas ascendentemente', 50)
    
    ciudades_des = list(sorted(ciudades, reverse=True))
    mostrar_ciudades(ciudades_des, 'ciudades ordenadas desendentemente', 50)
    
    
def mostrar_ciudades(ciudades: List[str], titulo: str, largo: int) -> None:
    encabezado(titulo, largo)
    for item in ciudades:
        print(f"la ciudad es {item.upper()}")
    
main()