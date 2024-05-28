package main

import (
	"fmt"
)

func leerNota() float32 {
	var nota float32
	fmt.Print("Ingresa una nota: ")
	fmt.Scan(&nota)
	return nota
}

func calcularPromedio(notas []float32) float32 {
	suma := float32(0)
	for _, nota := range notas {
		suma += nota
	}
	return suma / float32(len(notas))
}

func main() {
	var cantEstudiantes int
	fmt.Print("Ingresa la cantidad de estudiantes: ")
	fmt.Scan(&cantEstudiantes)

	notasClase := make([]float32, cantEstudiantes)
	mayorNota := float32(0)
	menorNota := float32(1)

	for i := 0; i < cantEstudiantes; i++ {
		fmt.Printf("Estudiante %d:\n", i+1)
		notas := make([]float32, 4)
		for j := 0; j < 4; j++ {
			notas[j] = leerNota()
		}
		promedioEstudiante := calcularPromedio(notas)
		notasClase[i] = promedioEstudiante

		fmt.Printf("Prmedio %d: %.2f\n", i+1, promedioEstudiante)

		for _, nota := range notas {
			if nota > mayorNota {
				mayorNota = nota
			}
			if nota < menorNota {
				menorNota = nota
			}
		}
	}

	promedioClase := calcularPromedio(notasClase)

	fmt.Printf("\nEl promedio de la clase fue de: %.2f\n", promedioClase)
	fmt.Printf("la nota más alta: %.2f\n", mayorNota)
	fmt.Printf("la nota más alta: %.2f\n", menorNota)
}
