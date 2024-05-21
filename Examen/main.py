import os

from utilidades import utilidades

r = None

lista_inventario={}
inventario = {}

while r == None:
    utilidades.clear()

    print("Inventario\n=============\n1- Agregar Producto\n2- Actualizar Productos\n3- Mostrar Productos\n4- Buscar Producto\n5- Salir\n")

    r = input("¿Que deseas hacer?: ")

    if r == str(1):
        utilidades.clear()

        n = input("Ingrese el nombre de su producto: ")
        c = input("ingrese la cantidad del producto: ")
        p = input("Ingrse el precio del producto: ")
        id = input("Ingrese la id del producto: ")
        inventario ={
            "Nombre": n,
            "Cantidad": c,
            "Precio":p,
            "id": id
        }
        lista_inventario[id] = inventario
        utilidades.clear()
        r = None
        continue

    if r == str(2):
        utilidades.clear()
        
        np = input("Ingrese el ID del producto que desea actualizar: ")
        
        if np in lista_inventario.keys():
            nc = input("Ingrese la nueva cantidad: ")
            inventario["Cantidad"] = nc
            
        print("La cantidad del producto fue actualizad con exito")
        input("Presione enter para continuar.......")
        r = None
        continue


    if r == str(3):
        utilidades.clear()
        print("Todos los productos disponibles:\n")
        for x, y in lista_inventario.items():
            print(f"{y}")
        input("Presione enter para continuar.......")
        r = None
        continue
    

    if r == str(4):
        utilidades.clear()
        h = input("Ingrese el ID del producto: ")
        lista_inventario.get(h, "El producto no existe en el inventario")
        print(f"{lista_inventario[h]}")
        input("Presione enter para continuar.......")
        r = None
        continue
           

    if r == str(5):
        input("Saliendo de la aplicacion......")
        exit()

