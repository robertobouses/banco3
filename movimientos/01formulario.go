package movimientos

import "fmt"

type DatosPersonales struct {
	dni      int
	nombre   string
	apellido string
}

type Movimiento struct {
	tipo    string
	importe int
	saldo   int
}

type CuentaBancaria struct {
	cliente   DatosPersonales
	gestiones Movimiento
	saldo     int
}

func (d *DatosPersonales) FormularioDatosPersonales() DatosPersonales {

	fmt.Println("Formulario de inscripción de CLIENTE")
	fmt.Println("Indica el dni")
	fmt.Scanln(&d.dni)
	fmt.Println("Indica el nombre")
	fmt.Scanln(&d.nombre)
	fmt.Println("Indica el apellido")
	fmt.Scanln(&d.apellido)

	Cliente1 := DatosPersonales{
		dni:      d.dni,
		nombre:   d.nombre,
		apellido: d.apellido,
	}
	return Cliente1
}

func GenerarMapa(Cliente1 DatosPersonales) {
	var InventarioCuentaBancaria map[int]CuentaBancaria
	InventarioCuentaBancaria = make(map[int]CuentaBancaria)
	var id int
Cuenta1:=CuentaBancaria{
	cliente
	
}

	InventarioCuentaBancaria[id] = Cliente1

}
