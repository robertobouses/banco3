package main

import (
	"fmt"

	"github.com/robertobouses/banco3/movimientos"
)

func main() {
	valor := movimientos.CuentaBancaria{}
	dato := movimientos.DatosPersonales{}
	fmt.Println("")
	dato.FormularioDatosPersonales()
	valor.Cargo(3)
	valor.Abono(1)

}
