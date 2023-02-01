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
	id        int
	cliente   DatosPersonales
	gestiones Movimiento
	saldo     int
}

func (c *CuentaBancaria) Cargo(monto int) {
	c.saldo += monto
	fmt.Printf("la cuenta %v de %v arroja un saldo de %v", c.id, c.cliente.nombre, c.saldo)
}
