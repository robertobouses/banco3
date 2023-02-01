package movimientos

import "fmt"

func (c *CuentaBancaria) Cargo(monto int) {
	c.saldo += monto
	fmt.Printf("la cuenta %v de %s arroja un saldo de %v\n", c.id, c.cliente.nombre, c.saldo)
}
