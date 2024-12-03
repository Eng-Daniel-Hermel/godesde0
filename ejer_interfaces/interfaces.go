package ejer_interfaces

import (
	"fmt"

	"github.com/Eng-Daniel-Hermel/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	if hu.Sexo() == "Hombre" {
		fmt.Printf("Soy un %s, y estoy respirando. \n", hu.Sexo())
	} else {
		fmt.Printf("Soy una %s, y estoy respirando. \n", hu.Sexo())
	}
}
