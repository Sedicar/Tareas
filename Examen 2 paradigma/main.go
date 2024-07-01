package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Song struct {
	Titulo  string
	Artista string
	Next    *Song
}

type Linkedlist struct {
	Head *Song
	Size uint32
}

func readerStrings(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Hubo un error, intente nuevamente", err)
		return ""
	}
	input = strings.TrimSpace(input)
	return input
}

func readerInt(message string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", message)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Hubo un error, intente nuevamente\n")
			continue
		}
		input = strings.TrimSpace(input)
		inputConv, errConv := strconv.Atoi(input)
		if errConv != nil {
			fmt.Print("Por Favor Ingrese un valor valido\n")
			continue
		}
		return inputConv
	}
}

func main() {
	lista := Linkedlist{}
	for {

		fmt.Println("      Reproductor\n", "==========================\n", "1-Agregar Cancion\n", "2-Eliminar Cancion\n", "3-Buscar Cancion\n", "4-Mostrar Canciones")
		opcion := readerInt("Escoje una opcion: ")

		switch opcion {
		case 1:
			titulo := readerStrings("Título de la canción: ")
			artista := readerStrings("Artista: ")
			lista.AgregarCancion(titulo, artista)
			fmt.Println("Canción añadida correctamente.")
		case 2:
			titulo := ("Título de la canción a eliminar: ")
			lista.BorrarCancion(titulo)
			fmt.Println("Canción eliminada correctamente.")
		case 3:
			titulo := ("Título de la canción a buscar: ")
			lista.BuscarCancion(titulo)
		case 4:
			fmt.Println("--- Todas las canciones ---")
			lista.ImprimirCanciones()
		case 5:
			fmt.Println("Chau")
			return
		}

	}

}

func (lista *Linkedlist) AgregarCancion(titulo, artista string) {
	newSong := &Song{Titulo: titulo, Artista: artista}
	if lista.Head == nil {
		lista.Head = newSong
	} else {
		current := lista.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newSong
	}
}

func (lista *Linkedlist) BorrarCancion(titulo string) {
	if lista.Head == nil {
		fmt.Println("No hay nada")
		return
	}

	if lista.Head.Titulo == titulo {
		lista.Head = lista.Head.Next
		return
	}

	current := lista.Head
	for current.Next != nil {
		if current.Next.Titulo == titulo {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
	fmt.Println("Canción no encontrada.")
}

func (lista *Linkedlist) BuscarCancion(titulo string) {
	current := lista.Head
	for current != nil {
		if current.Titulo == titulo {
			fmt.Printf("Canción encontrada:\nTítulo: %s\nArtista: %s\n", current.Titulo, current.Artista)
			return
		}
		current = current.Next
	}
	fmt.Println("Canción no encontrada.")
}

func (lista *Linkedlist) ImprimirCanciones() {
	current := lista.Head
	for current != nil {
		fmt.Printf("Título: %s, Artista: %s\n", current.Titulo, current.Artista)
		current = current.Next
	}
}
