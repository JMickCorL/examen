package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

var path = "test.txt"

type alumno struct {
	nombre        string
	calificacion1 int
	calificacion2 int
	calificacion3 int
}

var alum []alumno

func main() {
	var bandera = false

	for !bandera {
		fmt.Println("Que operación desea realizar?")
		fmt.Println("\n 1 = capturar \n 2 = guardar \n 3 = vizualizar informacion \n 4 = salir ")
		var operacion int
		fmt.Println("\nInserte la operacion: ")
		fmt.Scanln(&operacion)

		switch {
		case operacion == 1:
			alumnos()
		case operacion == 2:
			Write("test.txt", alum)
		case operacion == 3:
			leer()
		case operacion == 4:
			os.Exit(1)
		default:
			fmt.Println("Por favor, elija una operacion")
		}
	}
}

func alumnos() {

	for i := 0; i <= 4; i++ {

		a := alumno{}

		fmt.Println("\n Nombre: ")
		fmt.Scanln(&a.nombre)
		fmt.Println("\n calificacion 1: ")
		fmt.Scanln(&a.calificacion1)
		fmt.Println("\n calificacion 2: ")
		fmt.Scanln(&a.calificacion2)
		fmt.Println("\n calificacion 3: ")
		fmt.Scanln(&a.calificacion3)

		alum = append(alum, a)

	}

	fmt.Println(alum)
}

// Write will print any string of text to a file safely by
func Write(filename string, alumnos []alumno) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := ""
	for i := 0; i < len(alumnos); i++ {
		data += alumnos[i].nombre + strconv.Itoa(alumnos[i].calificacion1) + strconv.Itoa(alumnos[i].calificacion2) + strconv.Itoa(alumnos[i].calificacion3) + "\n"
	}
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func leer() {
	nombreArchivo := "test.txt"
	bytesLeidos, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}

	contenido := string(bytesLeidos)
	fmt.Printf("El contenido del archivo es: %s", contenido)
}
