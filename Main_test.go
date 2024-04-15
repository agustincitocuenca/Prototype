package main_test

import (
	"testing"

	clases "github.com/agustincitocuenca/Prototype/Clases"
)

func Test_ReporteMedico_Clonacion(t *testing.T) {
	reporteOriginal := clases.Constructor_ReporteMedico(
		"PDF",
		"Lucas",
		"Informe Mensual",
		"Este es el cuerpo del informe medico",
		"Dengue",
	)

	copiaReporte := (reporteOriginal.Clonar()).(*clases.ReporteMedico)

	SonIguales := &reporteOriginal == &copiaReporte

	cuerpo := reporteOriginal.Cuerpo == copiaReporte.Cuerpo
	formato := reporteOriginal.Formato == copiaReporte.Formato
	paciente := reporteOriginal.Paciente == copiaReporte.Paciente
	titulo := reporteOriginal.Titulo == copiaReporte.Titulo
	sintomas := reporteOriginal.Sintomas == copiaReporte.Sintomas
	losContenidosNoSonIguales := !(cuerpo && formato && paciente && titulo && sintomas)

	if SonIguales {
		t.Errorf("No se creo un clon del objeto, es el mismo objeto")
	}

	if losContenidosNoSonIguales {
		t.Errorf("los contenidos no son iguales")
		t.Errorf("cuerpo: %t", cuerpo)
		t.Errorf("formato: %t", formato)
		t.Errorf("paciente: %t", paciente)
		t.Errorf("titulo: %t", titulo)
		t.Errorf("sintomas: %t", sintomas)
	}

}
