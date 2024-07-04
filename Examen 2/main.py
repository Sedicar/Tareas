import pyodbc
import os
from src import conexion, funciones
      
while True:
    SERVER = 'DESKTOP-LV6EIRR\SQLEXPRESS'
    DATABASE = 'Escuela'
    TABLE = 'Cursos'
    conexion.CONEXION(SERVER, DATABASE)
    
    cursor = conexion.cursor()
    
    opciones=int(input("¿Que quieres? \n1-Agregar un Curso \n2-Leer los Cursos \n3-Actualizar un Curso \n4-Borrar un Curso"))
    match opciones:
        case 1:
            nombre=str(input("Ingrese el nombre del curso: \n"))
            descripcion=str(input("Ingrese la descripcion del curso: \n"))
            profesor_id=1
            funciones.crear(conexion, nombre, descripcion, profesor_id)
            print("Curso Creado con exito")
        case 2:
            funciones.leer(conexion)
        case 3:
            nombre=str(input("Ingrese el nuevo nombre del curso: \n"))
            descripcion=str(input("Ingrese ls nueva descripcion del curso: \n"))
            funciones.actualizar(conexion, nombre, descripcion)
            print("Curso Actualizado con exito")
        case 4:
            funciones.leer(conexion)
            curso_id=str(input("Ingrese la id del curso a eliminar"))
            funciones.eliminar(conexion, curso_id)
            print("Curso Eliminado con exito")    
    
    
    