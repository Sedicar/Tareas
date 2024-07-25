from django.shortcuts import render, HttpResponse, redirect, get_object_or_404
from myapp.models import *
from .forms import *

def lista_autores(request):
    autores = Autor.objects.all()
    for i in autores:
        print(i)
    return render(request, 'autor/lista.html', {
        'autores' : autores
    })
    

def crear_autor(request):
    if request.method == 'POST':
        form = AutorForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('lista_autores')
    else:
        form = AutorForm()   
    return render(request, 'autor/crear.html', {
        'form': form
    })
    
def eliminar_autor(request, autor_id):
    autor = get_object_or_404(Autor, pk=autor_id)
    autor.delete()
    return redirect('lista_autores')
    
def editar_autor(request, autor_id):
    autor = get_object_or_404(Autor, pk=autor_id)
    if request.method == 'POST':
        form = AutorForm(request.POST, instance=autor)
        if form.is_valid():
            form.save()
            return redirect('lista_autores')
    else:
        form = AutorForm(instance=autor)
    return render(request, 'autor/editar_autor.html', {'form': form, 'autor': autor})
    