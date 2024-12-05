package main

import (
	"fmt"

	"github.com/Eng-Daniel-Hermel/godesde0/goroutines"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os=="linux" || os=="OS X." {
		fmt.Println("Esto no es Windows, es", os)
	} else {
		fmt.Println("Esto es", os)
	}

	switch os := runtime.GOOS; os  {
	case "linux":
		fmt.Println("Esto es", os)
	case "darwin":
		fmt.Println("Esto es", os)
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.ConvNumerico("ffffff")
	fmt.Println(numero)
	fmt.Println(texto

	teclado.IngresoNumeros()

	iteraciones.Iterar()

	fmt.Println(ejercicios.TabladeMultiplicar())*/

	//files.GravarTabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)
	//arreglosslices.MuestroArreglos()
	//arreglosslices.MuestroSlice()
	//arreglosslices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria:=new(modelos.Mujer)
	e.HumanosRespirando(Maria)*/

	//defer_panic.VemosDefer()

	//defer_panic.EjemploPanic()

	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Daniel Hermel", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")

}
