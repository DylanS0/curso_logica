'''generador de tablas de multiplicar'''

import utilidades as ut

def main():
    ut.linea(50)
    print("Bienvenido al generador de tablas de multiplicar".upper())
    num : int = 0
    inicio : int = 0
    final : int = 0
    
    ut.encabezado('generador de tablas', 50)
    num = ut.leerNumero('ingresa tu numero', int)
    ut.encabezado(f'tabla de multiplicar del {num}', 50)
    generarTabla(num)
    ut.encabezado('generar rango de tablas de multiplicar', 50)
    inicio = ut.leerNumero('ingresa el numero de inicio', int)
    final = ut.leerNumero('ingresa el numero de final', int)
    ut.encabezado(f'tabla de multiplicar del {inicio} al {final}', 50)
    generarRangoTablas(inicio, final)
    
def generarRangoTablas(inicio: int, final: int):
    nroTabla: int = abs( inicio - final ) + 1
    for i in range(nroTabla):
        for j in range (1, 11):
            print(f'{inicio} X {j} = {producto(inicio, j)}')
        ut.linea(50)
        inicio += 1
        
    
    
    
def generarTabla(num: int) -> None:
    for i in range(1,11):
        print(f'{num} x {i} = {producto(num, i)}')
        
def producto(a: int, b: int) -> int:
    return a * b
        
    
main()