package clases

import (
	"fmt"
	"math/rand"

	interfaces "github.com/agustincitocuenca/Prototype/Interfaces"
)

type ReporteMedico struct {
	Paciente         string
	Titulo           string
	Sintomas         string
	Cuerpo           string
	Formato          string
	numeroDelReporte int
}

func Constructor_ReporteMedico(
	formato string,
	paciente string,
	titulo string,
	cuerpo string,
	sintomas string) *ReporteMedico {
	numeroDelReporte := rand.Int()
	return &ReporteMedico{
		Formato:          formato,
		Paciente:         paciente,
		Titulo:           titulo,
		Cuerpo:           cuerpo,
		Sintomas:         sintomas,
		numeroDelReporte: numeroDelReporte,
	}
}

// Clonar crea y devuelve un clon del reporte medico.
func (t *ReporteMedico) Clonar() interfaces.IPrototype {
	return &ReporteMedico{
		Formato:          t.Formato,
		Titulo:           t.Titulo,
		Paciente:         t.Paciente,
		Sintomas:         t.Sintomas,
		Cuerpo:           t.Cuerpo,
		numeroDelReporte: t.numeroDelReporte,
	}
}

func (r *ReporteMedico) ImprimirReporte() {
	fmt.Println("Formato: " + r.Formato)
	fmt.Println("Titulo: " + r.Titulo)
	fmt.Println("Paciente: " + r.Paciente)
	fmt.Println("Sintomas: " + r.Sintomas)
	fmt.Println("Cuerpo: " + r.Cuerpo)

}
