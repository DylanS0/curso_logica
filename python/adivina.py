"""Juego adivina el numero"""
from utilidades import *
from random import randint

def main():
    #declarar variables
    secreto: int =  randint(1, 20)
    numero: int = 0
    alias: str = ""
    
    encabezado('juego de adivina el numero', 50)
    alias = leerTexto("ingresa tu alias", "alias")
    linea(50)
    jugar(secreto, numero, alias)
    
def jugar (secreto: int, numero: int, alias: str) -> None:
    texto: str = ''
    intentos : int = 1
    while numero != secreto:
        
        numero = leerNumero()
        intentos = intentos + 1
        
        
        if intentos == 4:
            
            texto = 'perdiste, el numero secreto era el'
            linea(50)
            print(f'{alias.upper()}, {texto} {secreto}')
            linea(50)
            break
        
        if numero > secreto:
            texto = 'tu numero es mayor que el secreto...'
            print(f'{alias.upper()}, {texto}')
            
        elif numero < secreto:
            texto = 'tu numero es menor que el secreto...'
            print(f'{alias.upper()}, {texto}')
            
        else:
            linea(50) 
            texto = '¡felicidades, adivinaste el numero secreto!'
            print(f'{alias.upper()}, {texto}')
            linea(50) 
        

main()