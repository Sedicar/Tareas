package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := ""
	Nombre := "sistema"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(localhost:3306)/"+Nombre)

	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)

	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)
	conexionEstablecida := conexionBD()
	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM consolas WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	borrarRegistro.Exec(idEmpleado)

	http.Redirect(w, r, "/", 301)
}

type consolas struct {
	Id      int
	Nombre  string
	Precio  string
	Detalle string
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conexionBD()
	registros, err := conexionEstablecida.Query("SELECT * FROM consolas")

	if err != nil {
		panic(err.Error())
	}

	empleado := consolas{}
	arregloEmpleado := []consolas{}

	for registros.Next() {
		var id int
		var nombre, precio, detalle string
		err = registros.Scan(&id, &nombre, &precio, &detalle)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Precio = precio
		empleado.Detalle = detalle

		arregloEmpleado = append(arregloEmpleado, empleado)
	}

	plantillas.ExecuteTemplate(w, "inicio", arregloEmpleado)
}

func Editar(w http.ResponseWriter, r *http.Request) {

	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)

	conexionEstablecida := conexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM consolas WHERE id =?", idEmpleado)

	empleado := consolas{}
	for registro.Next() {
		var id int
		var nombre, precio, detalle string
		err = registro.Scan(&id, &nombre, &precio, &detalle)

		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Precio = precio
		empleado.Detalle = detalle

	}
	fmt.Println(empleado)
	plantillas.ExecuteTemplate(w, "editar", empleado)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		precio := r.FormValue("precio")
		detalle := r.FormValue("detalle")
		conexionEstablecida := conexionBD()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO consolas(nombre,precio,detalle) VALUES(?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, precio, detalle)

		http.Redirect(w, r, "/", 301)
	}
}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		precio := r.FormValue("precio")
		detalle := r.FormValue("detalle")

		conexionEstablecida := conexionBD()
		modificarRegistros, err := conexionEstablecida.Prepare("UPDATE consolas SET nombre=?,precio=?,detalle=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}
		modificarRegistros.Exec(nombre, precio, detalle, id)

		http.Redirect(w, r, "/", 301)
	}
}
