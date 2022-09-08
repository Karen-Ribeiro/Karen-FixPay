package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso    string
	faculdade string
	semestre int
}

func main(){
	fmt.Println("Herança")

	p1 := pessoa{"Karen", "Ribeiro", 19, 167}
	e1 := estudante{p1, "Eng. de Computação", "IFCE", 3}

	fmt.Println(p1, e1)
}