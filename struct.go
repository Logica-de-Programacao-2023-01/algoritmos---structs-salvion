package main

//Crie uma struct chamada Circulo com o campo "raio". Escreva uma função que receba um Circulo como parâmetro e
//calcule a área do círculo (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi.
import "fmt"

func main() {
	var c Circulo
	fmt.Printf("qual o raio do circulo?\n")
	fmt.Scan(&c.raio)

	area := CalcArea(c)
	fmt.Print(area)
}

type Circulo struct {
	raio float64
}

func CalcArea(c Circulo) float64 {
	pi := 3.141592
	area := pi * c.raio * c.raio
	return area
}
