from django.urls import path
from myapp.views import *
from .views import eliminar_autor

urlpatterns = [
    path('autor/', lista_autores, name="lista_autores"),
    path('autor/crear/', crear_autor, name="crear_autor"),
    path('eliminar_autor/<int:autor_id>/', eliminar_autor, name='eliminar_autor'),
    path('editar_autor/<int:autor_id>/', editar_autor, name='editar_autor'),
]