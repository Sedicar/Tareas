import pyodbc

def CONEXION(SERVER, DATABASE):
    try:
        conexion = pyodbc.connect(f'DRIVER={{ODBC Driver 17 for SQL SERVER}};SERVER={SERVER};DATABASE={DATABASE};Trusted_Conecction=yes;')
        print("Conexion epica")
    except Exception as ex:
        print(f'Error al conectar la base de datos {ex}')   