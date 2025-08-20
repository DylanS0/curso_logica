"""
Primer programa en python
"""
#Programa que suma dos numeros enteros
import utilidades as ut

def main():
    numero1: int = 5
    numero2: int = 5
    suma: int = numero1 + numero2
    
    ut.encabezado("sumar numeros ", 50)
    
    print(f"La suma de {numero1} y {numero2} es {suma}".center(50))
    ut.linea(50)
    
main()
