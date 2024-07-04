import pyodbc

def crear(conexion, nombre, descripcion, profesor_id):
  try:
    cursor = conexion.cursor()
    valores = (nombre, descripcion, profesor_id)
    cursor.execute(valores)
    conexion.commit()
    return True
  except Exception as e:
    print(f"Error al crear curso: {e}")
    return False

def leer(conexion):
  try:
    cursor = conexion.cursor()
    resultados = cursor.fetchall()
    return resultados
  except Exception as e:
    print(f"Error al leer cursos: {e}")
    return None

def actualizar(conexion, nombre, descripcion):
  try:
    cursor = conexion.cursor()
    valores = (nombre, descripcion)
    cursor.execute(valores)
    conexion.commit()
    return True
  except Exception as e:
    print(f"Error al actualizar curso: {e}")
    return False

def eliminar(conexion, curso_id):
  try:
    cursor = conexion.cursor()
    valor = (curso_id,)
    cursor.execute(valor)
    conexion.commit()
    return True
  except Exception as e:
    print(f"Error al eliminar curso: {e}")
    return False
