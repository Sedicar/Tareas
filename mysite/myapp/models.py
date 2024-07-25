from django.db import models

class Autor(models.Model):
    name = models.CharField(max_length=100)
    
    
    class Meta:
        db_table = 'autores'
    
    def __str__(self):
        return self.name

class Libro(models.Model):
    title = models.CharField(max_length=100)
    autor = models.ForeignKey(Autor, on_delete=models.CASCADE, related_name='libros')
    
    class Meta:
        db_table = 'libros'
        
    def __str__(self):
        return self.title    
    
 
 