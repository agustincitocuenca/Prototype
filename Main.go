package main

import (
	"fmt"

	clases "github.com/agustincitocuenca/Prototype/Clases"
)

func main() {
	reporteOriginal := clases.Constructor_ReporteMedico(
		"PDF",
		"Lucas",
		"Informe Mensual",
		"Este es el cuerpo del informe medico",
		"Dengue",
	)

	nuevoReporte := (reporteOriginal.Clonar()).(*clases.ReporteMedico)
	nuevoReporte.Paciente = "Stiven"
	// Se crea un clon del objeto para aprovechar la configuracion anterior ya que en lo
	// unico que difieren es en el nombre del paciente ahorrandose todo el trabajo de crear uno casi igual.
	// Ademas de que las variables privadas no se podian conocer desde afuera asi que el mismo objeto es el que
	// se tenia que encargar de su clonacion

	fmt.Println("Reporte original:")
	reporteOriginal.ImprimirReporte()

	fmt.Println("Nuevo reporte:")
	nuevoReporte.ImprimirReporte()
}
