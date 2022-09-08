package main

import "fmt"

type Endereco struct{
	Rua string
	NumeroDaCasa int
	Bairro string
}

func  main(){
	cadastrarEndereco()

}

func cadastrarEndereco() {
	fmt.Println("Qual o nome da sua rua?")
	fmt.Scan(&e.Rua)
	fmt.Println("Qual o número da sua casa?")
	fmt.Scan(&e.NumeroDaCasa)
	fmt.Println("Qual o nome do seu bairro?")
	fmt.Scan(&e.Bairro)

	fmt.Println("Rua:", e.Rua, ",", e.NumeroDaCasa, ". ", e.Bairro )

}