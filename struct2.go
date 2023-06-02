package main

//Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço". O campo "endereço" deve ser uma outra
//struct com os campos "rua", "número", "cidade" e "estado".
//Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo.
import "fmt"

func main() {
	var p Pessoa
	fmt.Print("\nNome:")
	fmt.Scanf("&p.nome")
	fmt.Print("\nIdade:")
	fmt.Scanf("&p.idade")
	fmt.Print("\nEstado:")
	fmt.Scanf("&p.endereco.estado")
	fmt.Print("\nCidade:")
	fmt.Scanf("&p.endereco.cidade")
	fmt.Print("\nRua:")
	fmt.Scanf("&p.endereco.rua")
	fmt.Print("\nNúmero:")
	fmt.Scanf("&p.endereco.numero")
	OutputP(p)
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

type Endereco struct {
	rua    int
	numero int
	cidade string
	estado string
}

func OutputP(p Pessoa) {
	fmt.Printf(p.nome, p.endereco.estado, p.endereco.cidade, p.endereco)
}
