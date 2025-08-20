from typing import Any
def linea(largo: int):
    print("═" * largo) 
    
def encabezado(titulo: str, largo: int):
    linea(largo)
    print(titulo.upper().center(largo))
    linea(largo)
    
def leerNumero(msj:str, fn: Any) -> Any:
      numero: Any = 0
      while True:
          try: 
              numero = fn(input(msj + ': '))
              return numero
          except ValueError:
              print('debe escribir un numero...')
  
def leerTexto(msj: str, aviso: str) -> str:
    texto: str = ''
    while True:
        texto = input(msj + ": ")
        if texto.strip():
            return texto.strip()
        else: 
            print(f'debes escribir el {aviso}... ')
          
                    