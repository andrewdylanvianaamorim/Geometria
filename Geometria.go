package geometria

//Forma - é uma interface que deve ser implementado em qualquer forma geometrica
type Forma interface {
	CalculaArea() float64
}

//Retangulo - representa um retângulo e implementa Forma
type Retangulo struct {
	Largura float64
	Altura  float64
}

//CalculaArea - é a implementação de Forma
func (r Retangulo) CalculaArea() float64 {
	return r.Largura * r.Altura
}
