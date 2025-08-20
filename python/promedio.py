'''calcular el promedio de estatura de 6 personas y mostrar la estatura maxima y minima '''
from utilidades import *

def main():
    altura: float = 0
    altMax: float = 0
    altMin: float = 0
    nroPersona: int = 6
    promedio: float = 0
    sumatoria: float = 0   
    
    encabezado('ingreso de alturas', 50)
    
    for i in range(nroPersona):
        altura = leerNumero(f"Ingresa la altura {i + 1} en cm", float)
        sumatoria += altura
        
        if i == 0:
            altMax = altura
            altMin = altura
            
        if altura > altMax:
            altMax = altura
            
        if altura < altMin:
            altMin = altura
        
    promedio = sumatoria / nroPersona
    
    linea(50)
    print(f"Promedio de alturas: {promedio:.2f} cm")
    linea(50)
    print(f"Altura maxima: {altMax:.2f} cm")
    linea(50)
    print(f"Altura minima: {altMin:.2f} cm")

main()
