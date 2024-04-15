package interfaces

type IPrototype interface {
	Clonar() IPrototype
	ImprimirReporte()
}
