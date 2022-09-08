package main

import (
	"fmt"
	"time"
)

func mostraImc(imc float32){

	if imc >= 20 && imc <= 24{
		fmt.Println("peso normal")
	} else if imc >= 25 && imc <= 29{
		fmt.Println("excesso de peso")
	} else if imc >= 30 && imc <= 35{
		fmt.Println("obesidade")
	} else if imc >= 35{
		fmt.Println("super obesidade")
	}else if imc < 20{
		fmt.Println("peso abaixo do normal")
	}
}

func calculoImc(nome string, idade int, peso float32, altura float32) {

	fmt.Println("Qual o seu nome?")
	fmt.Scan(&nome)

	fmt.Println("Qual sua idade?")
	fmt.Scan(&idade)

	fmt.Println("Qual seu peso?")
	fmt.Scan(&peso)

	fmt.Println("Qual sua altura?")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Println("Bom dia,", nome, "! \n/Fazendo o calculo do Imc, eu vejo que você está com")

	mostraImc(imc)
}

func main() {

	var nome string
	var idade int
	var peso float32
	var altura float32

	fmt.Println("Hora:", time.Now())

	calculoImc(nome, idade, peso, altura)
}