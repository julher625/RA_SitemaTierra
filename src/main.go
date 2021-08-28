package main

import (
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	// Set up options.
	options := serial.OpenOptions{
		PortName:        "COM3",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	// Make sure to close it later.
	defer port.Close()

	b := make([]byte, 5) // Tamaño de la lectura
	exit := false
	for !exit {
		n, err := port.Read(b) //Leemos el puerto
		if err != nil {
			log.Fatalf("port.Read: %v", err)
		}
		//if n == len(b) { // Verificamos que tenga la longitud correcta
		fmt.Println("Read", n, "bytes.") //Cantidad de bytes leidos (7)
		fmt.Println(string(b))           //imprimimos la lectura
		//exit = true                      //Salimos del bucle
		//}
		var enviar = []byte("Hola")

		_, err = port.Write(enviar)
		if err != nil {
			log.Fatalf("port.Read: %v", err)
		}
	}

}
