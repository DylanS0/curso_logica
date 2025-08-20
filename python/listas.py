"""listas en python
"""
from utilidades import encabezado, linea
from typing import List

def main():
    
    #declarar una lista de edades
    edades: List[int] = [21,17,15,25,27,21,20, 18, 14,17]
    linea(50)
    print(f'Cantidad de edades es {len(edades)}')
    linea(50)
    print(f'La cuarta edad es {edades[3]}')
    linea(50)
    print(edades[1:4])
    linea(50)
    print(f"La ultima edad es {edades[-1]}")
    linea(50)
    print(f'La antepenultima edad es {edades[-2]}')
    encabezado('leer listas de edades', 50)
    for edad in edades:
        print(f'la edad es {edad}')
    linea(50)
    
    encabezado('leer lista de edades con indices', 50)
    for i, edad in enumerate(edades):
        print(f'la edad en la posicion {i} es {edad}')
        
    encabezado('leer lista de edades con indices al reves', 50)
    for edad in reversed(edades):
        print(f'la edad es {edad}')
    
    encabezado('edades ordenadas ascendente ', 50)
    edades.sort()
    for edad in edades:
        print(f'la edad es {edad}')
        
    encabezado('edades ordenadas desendente ', 50)
    edades.sort(reverse=True) 
    for edad in edades:
        print(f'la edad es {edad}')
        
    encabezado('mayores de edad', 50)
    mayores_edad = list(filter(lambda edad: edad >=18, edades))
    for edad in mayores_edad:
        print(f'la edad es {edad}')
    
    encabezado('menores de edad', 50)
    menores_edad = list(filter(lambda edad: edad <=18, edades))
    for edad in menores_edad:
        print(f'la edad es {edad}')
        
    encabezado('edades unicas', 50)
    edades_unicas = list(set(edades))
    for edad  in edades_unicas:
        print(f"la edad es {edad}")  
        
        
        
        
    
main()