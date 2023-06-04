package main

import (
	"fmt"
	"net/http"
)

// Las funciones que sirven como handler toman un http.ResponseWriter y un http.Request como argumentos.
// El ResponseWriter se utiliza para devolver la respuesta HTTP.
func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hola soy Ceci :) \n")
}

func main() {
	// Registramos nuestros handlers en las rutas del servidor utilizando la función http.HandleFunc.
	// Configuramos el router predeterminado en el paquete net/http y tomamos una función como argumento.
	// Finalmente, llamamos a ListenAndServe con el puerto “:8080” y un controlador.
	// El nil le dice que use el router predeterminado que acabamos de configurar.
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}
