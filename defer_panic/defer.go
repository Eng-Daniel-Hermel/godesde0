package defer_panic

import (
	"fmt"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mesaje final")

	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
