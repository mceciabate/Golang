package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// EJEMPLO ERRORES
func main() {

	//CON NEW()
	statusCode := 404
	if statusCode > 399 {
		fmt.Println(errors.New("La petición ha fallado"))
		return
	}
	fmt.Println("El programa finalizó correctamente")

	//UNWRAP()
	error_1 := fmt.Errorf("first error")
	error_2 := fmt.Errorf("second error: %w", error_1)

	err := errors.Unwrap(error_2)
	if err == error_1 {
		fmt.Println("same error")
	}

	//IS()
	_, e := os.Open("not_exist.txt")
	var notExist error = fs.ErrNotExist
	if errors.Is(e, notExist) {
		fmt.Println("The file does not exist")
	}

	//AS()
	_, e2 := os.Open("not_exist.txt")
	var pathError *fs.PathError
	if errors.As(e2, &pathError) {
		fmt.Println("The file does not exist", pathError.Path)
	}
	//PANIC
	fmt.Println("Iniciando...")
	_, e3 := os.Open("no-file.txt")
	if e3 != nil {
		panic(e3)
	}
	fmt.Println("Fin")
}
